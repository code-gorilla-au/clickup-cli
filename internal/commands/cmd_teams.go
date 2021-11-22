package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	teamsURL = fmt.Sprintf("%s/v2/team", baseURL)
)

func teamsCmd(fetchSvc fetchClient) *cobra.Command {
	return &cobra.Command{
		Use:   "teams",
		Short: "a new command",
		Long:  "a new command",
		RunE:  teamsFunc(fetchSvc),
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
		defer resp.Body.Close()
		var respBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
			log.Println("Error decoding response: ", err)
			return err
		}
		return nil
	}
}
