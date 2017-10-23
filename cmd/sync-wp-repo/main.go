package main

import (
	"fmt"
	"github.com/xwp/go-tide/wp"
	"time"
	"sync"
	"github.com/nanobox-io/golang-scribble"
	"github.com/xwp/go-tide/queue"
	"os"
)

// workerFetchFromRepo will periodically run through all the plugins in the WordPress.org repository
func workerFetchFromRepo(plugins chan repo.Plugin, perPage int) {

	currentPage := 0
	totalPages := 1

	fmt.Println("Getting plugin data from WordPress.org repo.")

	// Forever loop with time.Sleep() to give the fetching a break.
	for {

		fmt.Print("\n Sync started.\n")

		for currentPage < totalPages {
			currentPage += 1

			repoPlugins, err := repo.RequestPlugins("updated", perPage, currentPage)

			if err != nil {
				return;
			}

			totalPages = repoPlugins.Info.Pages
			fmt.Printf(" [%d/%d] ", currentPage, totalPages)

			for _, plugin := range repoPlugins.Plugins {
				plugins <- plugin
			}
		}

		fmt.Print("\nDone.\nWaiting for next sync.\n")

		// Wait before we go again.
		time.Sleep(time.Minute * 10)
	}
}

// workerRequestTideAudit pops new plugins from the plugins channel for processing.
func workerRequestTideAudit(plugins chan repo.Plugin, db *scribble.Driver) {

	awsRegion := getEnv("AWS_SQS_REGION", "")
	awsKey := getEnv("AWS_SQS_KEY", "")
	awsSecret := getEnv("AWS_SQS_SECRET", "")
	awsQueue := getEnv("AWS_SQS_QUEUE_NAME", "")

	sqsMan := queue.New(awsRegion, awsKey, awsSecret, awsQueue)

	if *sqsMan.QueueUrl == "" {
		panic(fmt.Sprintf("Fatal Error: Could not get Queue!" ))
	}

	for {
		select {
		case plugin := <-plugins:
			processPlugin(plugin, db, sqsMan)
		}
	}
}

// processPlugin sends audit request to API.
func processPlugin(plugin repo.Plugin, db *scribble.Driver, queueManager *queue.SqsManager) error {
	record := repo.Plugin{} // Record schema for db.Read().

	// Read current entry for plugin.
	db.Read("plugins", plugin.Slug, &record)
	if record.LastUpdated != plugin.LastUpdated || record.Version != plugin.Version {

		err := queueManager.SendToSQS(plugin.DownloadLink, "plugin")

		if err != nil {
			return err
		}

		if err := db.Write("plugins", plugin.Slug, plugin); err != nil {
			return err
		}

		fmt.Print("+")
	} else {
		fmt.Print(".")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Wait group to block main()
	wg := sync.WaitGroup{}
	wg.Add(1)

	perPage := 250;
	apiWorkers := perPage; // Each go routine consumes approx 4.5kb memory.

	//Setup the flat json db
	dbPath := getEnv("GO_TIDE_DATA", "./db")
	db, _ := scribble.New(dbPath, nil)

	// Plugins channel
	plugins := make(chan repo.Plugin)

	// Start worker to fetch plugins from repo.
	go workerFetchFromRepo(plugins, perPage);

	// Start multiple workers to add plugins to Tide.
	for i := 1; i <= apiWorkers; i++ {
		go workerRequestTideAudit(plugins, db)
	}

	// Wait forever.
	wg.Wait()
}
