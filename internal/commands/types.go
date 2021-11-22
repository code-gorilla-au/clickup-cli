package commands

import "github.com/spf13/cobra"

type Config struct {
	Token              string `json:"token,omitempty"`
	DefaultWorkspaceID string
}

type cmdWithErrorFunc = func(cmd *cobra.Command, args []string) error
