package commands

import (
	"log"

	"github.com/code-gorilla-au/clickup-cli/internal/store"
	"github.com/spf13/cobra"
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
	config, err := loadConfig(storeSvc)
	if err != nil {
		log.Println("Error loading config ", err)
		return err
	}
	configCmd := configCmd(storeSvc)
	configCmd.PersistentFlags().StringVarP(&tokenFlag, "token", "t", "", "your personal access token")
	configCmd.PersistentFlags().StringVarP(&teamIDFlag, "team-id", "i", "", "your workspace / team id")

	// add commands
	rootCmd.AddCommand(spacesCmd(clickSvc))
	rootCmd.AddCommand(teamsCmd(storeSvc, clickSvc, config))
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd())
	return rootCmd.Execute()
}

func loadConfig(storeSvc storage) (Config, error) {
	var config Config
	if err := storeSvc.Load(&config); err != nil {
		return config, err
	}
	return config, nil
}
