package commands

import (
	"github.com/spf13/cobra"
)

var (
	// root vars
	dryrunFlag  = false
	versionFlag = "v.dev-1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "clickup [COMMAND] --[FLAGS]",
	Short: "simple click up cli",
	Long:  `simple click up cli`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryrunFlag, "dry-run", "d", false, "run the cli in dry run mode")
}

func Execute() error {
	// add commands
	rootCmd.AddCommand(configCmd())
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}
