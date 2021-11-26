package commands

import (
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/formats"
	"github.com/spf13/cobra"
)

func teamsCmd(storeSvc storage, clickSvc clickupClient, config Config) *cobra.Command {
	return &cobra.Command{
		Use:   "teams",
		Short: "Get all of the user's workspace (teams)",
		Long: `
		Teams is the legacy term for what are now called Workspaces in ClickUp. 
		For compatibility, the term team is still used in this API. 
		This is not the new "Teams" feature which represents a group of users.`,
		PreRunE: checkPreReqConfig(config),
		RunE:    teamsFunc(clickSvc, config),
	}
}

func teamsFunc(clickSvc clickupClient, config Config) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		resp, err := clickSvc.GetAllWorkspaces(config.Token)
		if err != nil {
			log.Println("Error getting teams: ", err)
			return err
		}
		formats.PrintJSON(resp)
		return nil
	}
}
