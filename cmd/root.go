package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gvm",
		Short: "A golang version manager",
		Long:  `A Simple dependency free golang version manager`,
	}

	GoInstallationDirectory string
	GoPathFile              string
)

func init() {

	HomeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	GoInstallationDirectory = fmt.Sprintf("%s/.gvm", HomeDirectory)
	GoPathFile = fmt.Sprintf("%s/go_path", GoInstallationDirectory)

	if _, err := os.Stat(GoInstallationDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(GoInstallationDirectory, 0775); err != nil {
			panic(err)
		}
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
