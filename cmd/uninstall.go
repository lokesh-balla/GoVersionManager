package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// uninstallCmd represents the uninstall command.
var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"remove"},
	Short:   "Specify the version of golang needed to uninstalled/removed",
	Long: `Specify the version of golang needed to uninstalled/removed

		# To remove a specific version
		$ gvm uninstall 1.19
	`,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			if err := removeGoVersion(args[0]); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func removeGoVersion(version string) error {
	// check if valid installed version
	ok, err := checkVersionInstalled(version)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("specified version: %s is not installed", version)
	}

	ok, err = checkVersionDefault(version)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("present version of golang is set as default, please change it and try again")
	}

	if err := os.RemoveAll(fmt.Sprintf("%s/%s", GoInstallationDirectory, version)); err != nil {
		return err
	}

	return DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DBBucketName))
		if err != nil {
			return err
		}

		return bucket.Delete([]byte(version))
	})
}
