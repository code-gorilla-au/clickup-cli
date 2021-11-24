package commands

import (
	"github.com/code-gorilla-au/clickup-cli/internal/store"
	"github.com/spf13/cobra"
)

const (
	baseURL = "https://api.clickup.com/api"
)

var (
	commandsConfig Config
)

var (
	// root vars
	dryrunFlag  = false
	versionFlag = "v.dev-1.0.0"

	// config vars
	tokenFlag  string
	teamIDFlag string
)

var rootCmd = &cobra.Command{
	Use:   "clickup [COMMAND] --[FLAGS]",
	Short: "simple click up cli",
	Long:  `simple click up cli`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&dryrunFlag, "dry-run", "d", false, "run the cli in dry run mode")
}

func Execute(storeSvc *store.Service, clickSvc clickupClient) error {
	configCmd := configCmd(storeSvc)
	configCmd.PersistentFlags().StringVarP(&tokenFlag, "token", "t", "", "your personal access token")
	configCmd.PersistentFlags().StringVarP(&teamIDFlag, "team-id", "i", "", "your workspace / team id")

	// add commands
	rootCmd.AddCommand(spacesCmd(clickSvc))
	rootCmd.AddCommand(teamsCmd(storeSvc, clickSvc, commandsConfig))
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}

func beforeRunCmd(storeSvc storage) cmdWithErrorFunc {
	return func(cmd *cobra.Command, args []string) error {
		var config Config
		if err := storeSvc.Load(&config); err != nil {
			return err
		}
		commandsConfig = config
		return nil
	}
}
