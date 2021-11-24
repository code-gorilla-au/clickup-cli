package commands

import (
	"fmt"
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/formats"
	"github.com/spf13/cobra"
)

var (
	teamsURL = fmt.Sprintf("%s/v2/team", baseURL)
)

func teamsCmd(storeSvc storage, clickSvc clickupClient) *cobra.Command {
	return &cobra.Command{
		Use:   "teams",
		Short: "Get all of the user's workspace (teams)",
		Long: `
		Teams is the legacy term for what are now called Workspaces in ClickUp. 
		For compatibility, the term team is still used in this API. 
		This is not the new "Teams" feature which represents a group of users.`,
		PreRunE: beforeRunCmd(storeSvc),
		RunE:    teamsFunc(clickSvc),
	}
}

func teamsFunc(clickSvc clickupClient) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		resp, err := clickSvc.GetAllWorkspaces(personalToken)
		if err != nil {
			log.Println("Error getting teams: ", err)
			return err
		}
		formats.PrintJSON(resp)
		return nil
	}
}
