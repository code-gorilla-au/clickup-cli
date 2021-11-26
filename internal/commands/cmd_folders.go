package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func foldersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "folders",
		Short: "a new command",
		Long:  "a new command",
		Run:   foldersFunc,
	}
}

func foldersFunc(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		log.Println("at least 1 argument name must be provided")
		return
	}
}
