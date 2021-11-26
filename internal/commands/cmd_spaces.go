package commands

import (
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/formats"
	"github.com/spf13/cobra"
)

func spacesCmd(clickSvc clickupClient, config Config) *cobra.Command {
	return &cobra.Command{
		Use:     "spaces",
		Short:   "get all of a workspace's (teams) spaces",
		Long:    "get all of a workspace's (teams) spaces",
		PreRunE: checkPreReqConfig(config),
		RunE:    spacesFunc(clickSvc, config),
	}
}

func spacesFunc(clickSvc clickupClient, config Config) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		resp, err := clickSvc.GetSpaces(config.DefaultWorkspaceID, config.Token)
		if err != nil {
			log.Printf("Error getting spaces for workspace [%s]: %s", config.DefaultWorkspaceID, err)
			return err
		}
		formats.PrintJSON(resp)
		return nil
	}
}
