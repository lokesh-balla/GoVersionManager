package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Specify the version of golang needed to uninstalled/removed",
	Long: `Specify the version of golang needed to uninstalled/removed

		# To remove a specific version
		$ gvm uninstall 1.19
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if err := removeGoVersion(args[0]); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("only a single argument allowed for command uninstall")
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func removeGoVersion(version string) error {

	// check if valid installed version
	if !checkVersionInstalled(version) {
		return fmt.Errorf("specified version: %s is not installed", version)
	}

	if checkVersionDefault(version) {
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
