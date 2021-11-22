package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "prints out version of cli",
		Long:  "prints out version of cli",
		Run:   version,
	}
}

func version(cmd *cobra.Command, args []string) {
	log.Println(versionFlag)
}
