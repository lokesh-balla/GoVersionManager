package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/term"
)

const (
	DBBucketName = "METADATA"
	DEFAULT      = "default"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gvm",
		Short: "A golang version manager",
		Long:  `A Simple dependency free golang version manager`,
	}

	GoInstallationDirectory string
	GoPath              string

	DBPath string
	DB     *bolt.DB

	TerminalWidth  int
	TerminalHeight int
)

// Execute executes the root command.
func Execute() error {

	HomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	GoInstallationDirectory = fmt.Sprintf("%s/.gvm", HomeDirectory)
	if _, err := os.Stat(GoInstallationDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(GoInstallationDirectory, 0775); err != nil {
			return err
		}
	}
	GoPath = fmt.Sprintf("%s/go", GoInstallationDirectory)

	TerminalWidth, TerminalHeight, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	DBPath = fmt.Sprintf("%s/metadata.db", GoInstallationDirectory)

	DB, err = bolt.Open(DBPath, 0666, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return err
	}

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
