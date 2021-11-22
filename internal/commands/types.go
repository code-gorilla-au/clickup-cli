package commands

import "github.com/spf13/cobra"

type Config struct {
	Token string `json:"token,omitempty"`
}

type cmdWithErrorFunc = func(cmd *cobra.Command, args []string) error
