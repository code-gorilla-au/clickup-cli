package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/formats"
	"github.com/spf13/cobra"
)

var (
	teamsURL = fmt.Sprintf("%s/v2/team", baseURL)
)

func teamsCmd(storeSvc storage, fetchSvc fetchClient) *cobra.Command {
	return &cobra.Command{
		Use:   "teams",
		Short: "Get all of the user's workspace (teams)",
		Long: `
		Teams is the legacy term for what are now called Workspaces in ClickUp. 
		For compatibility, the term team is still used in this API. 
		This is not the new "Teams" feature which represents a group of users.`,
		PreRunE: beforeRunCmd(storeSvc),
		RunE:    teamsFunc(fetchSvc),
	}
}

func teamsFunc(fetchSvc fetchClient) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		resp, err := fetchSvc.Get(teamsURL, map[string]string{
			"Authorization": personalToken,
		})
		if err != nil {
			log.Println("Error getting teams: ", err)
			return err
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Println("Error closing body: ", err)
			}
		}()
		var respBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Println("Error decoding response: ", err)
			return err
		}
		formats.PrintJSON(&respBody)
		return nil
	}
}
