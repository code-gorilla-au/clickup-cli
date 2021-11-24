package commands

import "github.com/spf13/cobra"

type Config struct {
	Token              string `json:"token,omitempty"`
	DefaultWorkspaceID string `json:"default_workspace_id,omitempty"`
}

type cmdWithErrorFunc = func(cmd *cobra.Command, args []string) error
