package commands

import (
	"github.com/spf13/cobra"
)

func spacesCmd(clickSvc clickupClient) *cobra.Command {
	return &cobra.Command{
		Use:   "spaces",
		Short: "get all of a workspace's (teams) spaces",
		Long:  "get all of a workspace's (teams) spaces",
		RunE:  spacesFunc(clickSvc),
	}
}

func spacesFunc(clickSvc clickupClient) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		return nil
	}
}
