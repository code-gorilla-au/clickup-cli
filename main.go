package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/code-gorilla-au/clickup-cli/internal/chalk"
	"github.com/code-gorilla-au/clickup-cli/internal/commands"
)

const (
	envDev = "DEV"
)

var (
	cliEnv string
)

func init() {
	cliEnv, _ = os.LookupEnv("ENV")
}

func main() {
	setLogger()
	path, err := getStorePath()
	if err != nil {
		os.Exit(1)
	}
	if err := commands.Execute(path); err != nil {
		os.Exit(1)
	}
}

func setLogger() {
	log.SetPrefix(chalk.Cyan("clickup-cli "))
	if cliEnv == envDev {
		log.Println(chalk.Cyan("DEVELOPMENT MODE"))
		log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
		return
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

}

func getStorePath() (string, error) {
	if cliEnv == envDev {
		return "./config.json", nil
	}
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory: ", err)
		return "", err
	}
	flatPath := filepath.Join(dirname, ".clickup", "config.json")
	return flatPath, nil
}
