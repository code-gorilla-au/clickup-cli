package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func configCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Generates init config for click up",
		Long:  "Generates init config for click up",
		Run:   config,
	}
}

func config(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		log.Println("personal access token is required")
		return
	}
	config := Config{Token: args[0]}
	if storeService.Exists() {
		if err := storeService.Save(&config); err != nil {
			log.Println("Error saving token")
		}
		return
	}
	if err := storeService.Create(); err != nil {
		log.Println("Error creating config file")
		return
	}
	if err := storeService.Save(&config); err != nil {
		log.Println("Error saving token")
		return
	}
}
