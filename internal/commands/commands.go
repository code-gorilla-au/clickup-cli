package commands

import (
	"github.com/code-gorilla-au/clickup-cli/internal/store"
	"github.com/spf13/cobra"
)

var (
	personalToken string
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

func Execute(storeSvc *store.Service, fetchSvc client) error {
	// add commands
	rootCmd.AddCommand(configCmd(storeSvc))
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}

func beforeCmdRun(storeSvc storage) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		var config Config
		if err := storeSvc.Load(&config); err != nil {
			return err
		}
		personalToken = config.Token
		return nil
	}
}
