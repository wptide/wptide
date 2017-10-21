package main

import (
	"fmt"
	"github.com/xwp/go-tide/wp"
)

func main() {

	// Example Usage
	// @todo Remove this example
	fmt.Println("Getting plugin data from WordPress.org repo.")
	fmt.Println("============================================")
	repoPlugins, err := repo.RequestPlugins("updated", 5, 1)
	if err == nil {
		fmt.Printf("Current Page: %d \nTotal Pages: %d \nTotal Plugins: %d\n",
			repoPlugins.Info.Page,
			repoPlugins.Info.Pages,
			repoPlugins.Info.Results)
	}
	fmt.Println("Done.")
}
