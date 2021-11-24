package commands

import (
	"github.com/code-gorilla-au/clickup-cli/internal/store"
	"github.com/spf13/cobra"
)

const (
	baseURL = "https://api.clickup.com/api"
)

var (
	personalToken string
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

func Execute(storeSvc *store.Service, fetchSvc fetchClient) error {
	configCmd := configCmd(storeSvc)
	configCmd.PersistentFlags().StringVarP(&tokenFlag, "token", "t", "", "your personal access token")
	configCmd.PersistentFlags().StringVarP(&teamIDFlag, "team-id", "i", "", "your workspace / team id")

	// add commands
	rootCmd.AddCommand(teamsCmd(storeSvc, fetchSvc))
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
		personalToken = config.Token
		return nil
	}
}
