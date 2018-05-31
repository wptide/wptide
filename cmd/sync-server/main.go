package main

import (
	"github.com/wptide/pkg/sync"
	"github.com/wptide/pkg/message"
	"github.com/wptide/pkg/wporg"
	"time"
	"fmt"
	"strings"
	"github.com/wptide/pkg/env"
	"log"
	"strconv"
	"context"
	"github.com/wptide/pkg/sync/firestore"
)

var (
	// Sync Server config.
	conf            = getServiceConfig()
	serverActive    = conf["app"]["syncActive"] == "on"
	forcedSync      = conf["message"]["forceSync"] == "yes"
	apiClient       = conf["message"]["apiClient"]
	auditVisibility = conf["message"]["visibility"]
	perPage, _      = strconv.Atoi(conf["app"]["perPage"])
	poolSize, _     = strconv.Atoi(conf["app"]["poolSize"])
	poolDelay, _    = strconv.Atoi(conf["app"]["poolDelay"])
	browseCategory  = conf["app"]["browseCategory"]

	// Allow us to quit the server with a channel. Good for tests.
	quit = make(chan struct{}, 1)

	// Fetches themes and plugins from the WordPress.org API.
	client = &wporg.Client{}

	// Checks to see if we need to send an update.
	checker, checkerError = getSyncProvider(conf)

	// Create the SQS dispatcher.
	dispatcher, _ = getDispatcher(conf)

	// Tide API settings.
	apiVersion  = conf["tide"]["version"]
	apiProtocol = conf["tide"]["protocol"]
	apiHost     = conf["tide"]["host"]
	apiEndpoint = fmt.Sprintf("%s://%s/api/tide/%s/audit", apiProtocol, apiHost, apiVersion)
)

// fetcher returns a channel of RepoProjects to be processed by workers. It uses the
// "generator" concurrency pattern for feeding a channel.
func fetcher(projectType, category string, bufferSize int, token chan struct{}, maxPages int) <-chan wporg.RepoProject {

	// The new channel that will be processed by the workers.
	out := make(chan wporg.RepoProject, bufferSize)

	// Feed the channel concurrently.
	go func(token chan struct{}) {

		// Don't do anything until a token is available.
		<-token

		var response *wporg.ApiResponse
		var err error

		for {

			page := 1
			pages := 1

			// Get last sync start time for projectType, and subtract 5 minutes for a buffer.
			lastSync := checker.GetSyncTime("start", projectType).Add(-1 * (time.Minute * 5))

			// Set current sync time for projectType.
			checker.SetSyncTime("start", projectType, time.Now())

			for page <= pages {

				var timeFormat string
				if projectType == "themes" {
					timeFormat = wporg.TimeFormatThemes
					response, err = client.RequestThemes(category, bufferSize, page)
				} else {
					timeFormat = wporg.TimeFormat
					response, err = client.RequestPlugins(category, bufferSize, page)
				}

				// Normalize lastSync because of format differences.
				normalizedTime, _ := time.Parse(timeFormat, lastSync.Format(timeFormat))
				lastSync = normalizedTime

				if err != nil {
					log.Println(err)
					page = pages + 1
					continue
				}

				pages = response.Info.Pages
				if maxPages != -1 && pages > maxPages {
					pages = maxPages
				}

				// If we're doing a sync by "updated" items then only continue if
				// the most recently updated item is newer than our last sync time.
				mostRecent := mostRecentFromResponse(response)
				if category == "updated" && mostRecent != nil {
					// If the latest item's update time is less than the recorded last updated item
					// then set pages out of bounds to break the next loop.
					// We won't `continue` yet and might as well continue looping the items we have.
					if ! timeAfterEqual(*mostRecent, lastSync, timeFormat) {
						page = pages
						pages = 0
					}

				}

				// Themes.
				for _, project := range response.Themes {
					project.Type = projectType
					if category != "updated" || (category == "updated" && timeAfterEqual(project, lastSync, timeFormat)) {
						out <- project
					}
				}

				// Plugins.
				for _, project := range response.Plugins {
					project.Type = projectType
					if category != "updated" || (category == "updated" && timeAfterEqual(project, lastSync, timeFormat)) {
						out <- project
					}
				}

				page += 1
			}

			// Set sync stop time for projectType.
			checker.SetSyncTime("stop", projectType, time.Now())
			log.Println(projectType, "sync finished at", time.Now().Format(time.RFC3339))

			// Finished, so pass on the token.
			token <- struct{}{}

			time.Sleep(time.Second * time.Duration(poolDelay))

			// Wait until this process receives a token.
			<-token
		}
	}(token)

	return out
}

// pool starts a number of infoWorkers. This uses a "worker" pattern where each worker will
// read from a projects channel to process the results.
func pool(workers int, projects <-chan wporg.RepoProject, checker sync.UpdateChecker, dispatcher sync.Dispatcher, messages chan *message.Message) {
	for i := 0; i < workers; i++ {
		go infoWorker(projects, checker, dispatcher, messages)
	}
}

// infoWorker reads a project from the projects channel, runs it through the update check
// and (conditionally) sends it to the dispatcher which loads up the job queue.
func infoWorker(projects <-chan wporg.RepoProject, checker sync.UpdateChecker, dispatcher sync.Dispatcher, messages chan *message.Message) {
	for {
		select {
		case project := <-projects:
			// If this project is not in sync (or forced), then send it to the queue.
			if checker.UpdateCheck(project) || forcedSync {
				if err := dispatcher.Dispatch(project); err == nil {
					// If its been successfully added to the queue then change the
					// project state.
					checker.RecordUpdate(project)
				} else {
					log.Println(project.Name, err)
				}
			}
		}
	}
}

// defaultStandards gets the standards to run against a project.
func defaultStandards() []string {
	return []string{
		"phpcs_wordpress",
		"phpcs_phpcompatibility",
	}
}

// defaultAudits gets the audits to run against a project.
func defaultAudits() *[]message.Audit {
	return &[]message.Audit{
		{
			Type: "phpcs",
			Options: &message.AuditOption{
				Standard: "wordpress",
				Report:   "json",
				Encoding: "utf-8",
				Ignore:   "*/vendor/*,*/node_modules/*",
			},
		},
		{
			Type: "phpcs",
			Options: &message.AuditOption{
				Standard:   "phpcompatibility",
				Report:     "json",
				Encoding:   "utf-8",
				RuntimeSet: "testVersion 5.2-",
			},
		},
	}
}

// timeAfterEqual checks if the project's LastUpdated time is after the given time.
func timeAfterEqual(project wporg.RepoProject, t time.Time, format string) bool {
	pT, _ := time.Parse(format, project.LastUpdated)
	return pT.After(t) || t == pT
}

// mostRecentFromResponse returns the most recent project from the response.
func mostRecentFromResponse(response *wporg.ApiResponse) *wporg.RepoProject {
	var mostRecent *wporg.RepoProject

	if len(response.Themes) > 0 {
		mostRecent = &response.Themes[0]
	}

	if len(response.Plugins) > 0 {
		mostRecent = &response.Plugins[0]
	}

	return mostRecent
}

func main() {

	// SYNC_ACTIVE must be set to "on". This is to ensure you know what you are
	// doing and not flooding message queues.
	if ! serverActive {
		log.Println("Sync Server not active. Please set ENV variable `SYNC_ACTIVE=on`.")
		return;
	}

	// If sync client doesn't exist then exit.
	if checkerError != nil {
		log.Println("Sync client failed to initialize.")
		return;
	}

	// Message channel.
	messages := make(chan *message.Message, perPage)

	token := make(chan struct{}, 1)

	// themeTasks is a channel populated by a fetcher.
	themeTasks := fetcher("themes", browseCategory, perPage, token, -1)

	// Init the dispatcher.
	dispatcher.Init()

	// Start the theme worker pool.
	go pool(poolSize, themeTasks, checker, dispatcher, messages)

	// Send a token.
	token <- struct{}{}

	// Wait before starting the plugin fetcher, so that the theme can
	// pick up the token first.
	time.Sleep(time.Second * 2)

	// pluginTasks is a channel populated by a fetcher.
	pluginTasks := fetcher("plugins", browseCategory, perPage, token, -1)

	// Start the plugin worker pool.
	go pool(poolSize, pluginTasks, checker, dispatcher, messages)

	for {
		select {
		case <-quit:
			return
		}
	}
}

func getDispatcher(c map[string]map[string]string) (sync.Dispatcher, error) {
	switch c["app"]["messageProvider"] {
	case "sqs":
		conf := c["aws"]
		return &sqsDispatcher{
			Queues: map[string]struct {
				Region   string
				Key      string
				Secret   string
				Queue    string
				Endpoint string
				Active   bool
				Accepts  string // "all" or "themes" or "plugins"
			}{
				"phpcs": {
					conf["sqs_region"],
					conf["key"],
					conf["secret"],
					conf["sqs_phpcs_queue"],
					apiEndpoint,
					strings.ToLower(c["app"]["syncPHPCSActive"]) == "on",
					"all",
				},
				"lighthouse": {
					conf["sqs_region"],
					conf["key"],
					conf["secret"],
					conf["sqs_lh_queue"],
					apiEndpoint,
					strings.ToLower(c["app"]["syncLighthouseActive"]) == "on",
					"themes",
				},
			},
			providers: make(map[string]message.MessageProvider),
		}, nil
	case "firestore":
		conf := c["gcp"]

		return &firestoreDispatcher{
			ProjectID: conf["projectID"],
			Collections: map[string]struct {
				Collection string
				Active     bool
				Accepts    string // "all" or "themes" or "plugins"
			}{
				"phpcs": {
					conf["phpcsCollection"],
					strings.ToLower(c["app"]["syncPHPCSActive"]) == "on",
					"all",
				},
				"lighthouse": {
					conf["lighthouseCollection"],
					strings.ToLower(c["app"]["syncLighthouseActive"]) == "on",
					"themes",
				},
			},
			providers: make(map[string]message.MessageProvider),
		}, nil
	case "mongo":
		conf := c["mongo"]

		return &mongoDispatcher{
			ctx:      context.Background(),
			user:     conf["user"],
			pass:     conf["pass"],
			host:     conf["host"],
			database: conf["database"],

			Collections: map[string]struct {
				Collection string
				Active     bool
				Accepts    string // "all" or "themes" or "plugins"
			}{
				"phpcs": {
					conf["phpcsCollection"],
					strings.ToLower(c["app"]["syncPHPCSActive"]) == "on",
					"all",
				},
				"lighthouse": {
					conf["lighthouseCollection"],
					strings.ToLower(c["app"]["syncLighthouseActive"]) == "on",
					"themes",
				},
			},
			providers: make(map[string]message.MessageProvider),
		}, nil
	default:
		return nil, nil
	}
}

func getSyncProvider(c map[string]map[string]string) (sync.UpdateSyncChecker, error) {
	switch c["app"]["syncDBProvider"] {
	case "local":
		conf := c["local"]
		// Init the scribble (flat file) db used for checking currency of results.
		return &scribbleChecker{
			db: newScribbleChecker(conf["dbPath"]),
		}, nil
	case "firestore":
		conf := c["gcp"]
		syncClient, err := firestore.New(context.Background(), conf["projectID"], conf["docPath"])
		return syncClient, err
	default:
		return nil, nil
	}
}

func getServiceConfig() map[string]map[string]string {
	return map[string]map[string]string{
		"app": {
			"syncActive":           env.GetEnv("SYNC_ACTIVE", "off"),
			"syncPHPCSActive":      env.GetEnv("SYNC_PHPCS_ACTIVE", "on"),
			"syncLighthouseActive": env.GetEnv("SYNC_LH_ACTIVE", "off"),
			"perPage":              env.GetEnv("SYNC_ITEMS_PER_PAGE", "250"),
			"poolSize":             env.GetEnv("SYNC_POOL_WORKERS", "250"),
			"poolDelay":            env.GetEnv("SYNC_POOL_DELAY", "300"),
			"browseCategory":       env.GetEnv("SYNC_API_BROWSE_CATEGORY", "updated"),
			"syncDBProvider":       env.GetEnv("SYNC_DATABASE_PROVIDER", "local"),
			"messageProvider":      env.GetEnv("SYNC_MESSAGE_PROVIDER", "sqs"),
		},
		"message":
		{
			"forceSync":  strings.ToLower(env.GetEnv("SYNC_FORCE_AUDITS", "no")),
			"apiClient":  env.GetEnv("SYNC_DEFAULT_CLIENT", "wporg"),
			"visibility": env.GetEnv("SYNC_DEFAULT_VISIBILITY", "public"),
		},
		"local":
		{
			"dbPath": env.GetEnv("SYNC_DATA", "./db"),
		},
		"aws":
		{
			"key":             env.GetEnv("AWS_API_KEY", ""),
			"secret":          env.GetEnv("AWS_API_SECRET", ""),
			"sqs_region":      env.GetEnv("AWS_SQS_REGION", ""),
			"sqs_lh_queue":    env.GetEnv("AWS_SQS_QUEUE_LH", ""),
			"sqs_phpcs_queue": env.GetEnv("AWS_SQS_QUEUE_PHPCS", ""),
		},
		"gcp":
		{
			"projectID":            env.GetEnv("GCP_PROJECT", ""),
			"docPath":              env.GetEnv("SYNC_DATABASE_DOCUMENT_PATH", "sync-server/wporg"),
			"lighthouseCollection": env.GetEnv("GCF_QUEUE_LH", "queue-lighthouse"),
			"phpcsCollection":      env.GetEnv("GCF_QUEUE_PHPCS", "queue-phpcs"),
		},
		"mongo":
		{
			"user":                 env.GetEnv("MONGO_INITDB_ROOT_USERNAME", ""),
			"pass":                 env.GetEnv("MONGO_INITDB_ROOT_PASSWORD", ""),
			"host":                 env.GetEnv("MONGO_HOST", "localhost:27017"),
			"database":             env.GetEnv("MONGO_INITDB_DATABASE", "queue"),
			"lighthouseCollection": env.GetEnv("MONGO_QUEUE_LH", "lighthouse"),
			"phpcsCollection":      env.GetEnv("MONGO_QUEUE_PHPCS", "phpcs"),
		},
		"tide":
		{
			"host":     env.GetEnv("API_HTTP_HOST", "wptide.org"),
			"protocol": env.GetEnv("API_PROTOCOL", "https"),
			"version":  env.GetEnv("API_VERSION", "v1"),
		},
	}
}
