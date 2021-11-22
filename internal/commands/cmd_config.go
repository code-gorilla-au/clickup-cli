package commands

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func configCmd(ss storage) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Generates init config for click up",
		Long:  "Generates init config for click up",
		RunE:  configFunc(ss),
	}
}

func configFunc(ss storage) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Println("personal access token is required")
			return errors.New("no args")
		}
		config := Config{Token: args[0]}
		if ss.Exists() {
			if err := ss.Save(&config); err != nil {
				log.Println("Error saving token")
				return err
			}
		}
		if err := ss.Create(); err != nil {
			log.Println("Error creating config file")
			return err
		}
		if err := ss.Save(&config); err != nil {
			log.Println("Error saving token")
			return err
		}
		return nil
	}
}
