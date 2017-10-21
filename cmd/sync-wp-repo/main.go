package main

import (
	"fmt"
	"github.com/xwp/go-tide/wp"
)

func main() {

	// Example Usage
	// @todo Remove this example
	repoPlugins, err := repo.RequestPlugins("updated", 5, 1)
	if err == nil {
		fmt.Printf("Current Page: %d \nTotal Pages: %d \nTotal Plugins: %d",
			repoPlugins.Info.Page,
			repoPlugins.Info.Pages,
			repoPlugins.Info.Results)
	}
}
