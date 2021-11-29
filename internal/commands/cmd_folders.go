package commands

import (
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/formats"
	"github.com/spf13/cobra"
)

func foldersCmd(clickSvc clickUpClient, config Config) *cobra.Command {
	return &cobra.Command{
		Use:     "folders <space-id>",
		Short:   "Get all folders for a workspace",
		Long:    "Get all folders for a workspace",
		PreRunE: checkPreReqConfig(config),
		RunE:    foldersFunc(clickSvc, config),
	}
}

func foldersFunc(clickSvc clickUpClient, config Config) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			log.Println("space id is required")
			return ErrNoSpaceID
		}

		spaceID := args[0]

		folders, err := clickSvc.GetFolders(spaceID, config.Token)
		if err != nil {
			log.Println("Error getting workspace folders: ", err)
			return err
		}

		formats.PrintJSON(folders)
		return nil
	}
}
