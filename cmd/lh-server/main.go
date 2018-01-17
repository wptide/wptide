package main

import (
	"github.com/xwp/go-tide/src/audit"
	"github.com/xwp/go-tide/src/message"
	"github.com/xwp/go-tide/src/audit/ingest"
	"github.com/xwp/go-tide/src/audit/info"
	"github.com/xwp/go-tide/src/audit/lighthouse"
	"github.com/xwp/go-tide/src/audit/tide"
	"flag"
	"fmt"
	tideApi "github.com/xwp/go-tide/src/tide"
	"github.com/xwp/go-tide/src/tide/api"
)

var (
	Version    string // Set during build.
	Build      string // Set during build.
	TideClient tideApi.ClientInterface
)

func main() {

	// Is the -version flag being used?
	bVersion := flag.Bool("version", false, "a bool")

	// Parse all flags.
	flag.Parse()

	TideClient = &api.Client{}
	TideClient.Authenticate("","","")

	// @TODO: Code here.
	msg := message.Message{}
	processMessage(msg, TideClient)

	// If -version is requested then echo the version details.
	if *bVersion {
		fmt.Printf("Version: %s\nBuild: %s\n", Version, Build)
	}
}

// processMessage takes an SQS message and runs it through an audit process.
func processMessage(msg message.Message, client tideApi.ClientInterface) {

	// An slice of processes that need to be performed on the message.
	// A slice ensures that they happen in the correct order.
	processes := []audit.Processor{
		&ingest.Processor{},
		&info.Processor{},
		&lighthouse.Processor{},
		&tide.Processor{},
	}

	// Initialise an empty result.
	result := &audit.Result{
		"client": &client,
	}

	// Run through each processor and update the result.
	// Note: The result is a pointer so we're passing by reference.
	for _, proc := range processes {
		proc.Process(msg, result)
	}
}
