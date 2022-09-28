package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "specify the version of golang to use",
	Long: `specify the version of golang to use

		# To use go1.19 as the default
		$ gvm use go1.19
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if err := setGoVersion(args[0]); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("only a single argument allowed for command use")
		}
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func setGoVersion(version string) error {

	// check if valid installed version
	if !checkVersionInstalled(version) {
		return fmt.Errorf("specified version: %s is not installed", version)
	}

	// set PATH
	f, err := os.OpenFile(GoPathFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "export PATH=%s/%s/go/bin:$PATH", GoInstallationDirectory, version); err != nil {
		return err
	}

	// update DB
	return DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DBBucketName))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(version), []byte("default"))
	})
}
