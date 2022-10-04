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
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			if err := setGoVersion(args[0]); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func setGoVersion(version string) error {

	// check if valid installed version
	ok, err := checkVersionInstalled(version)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("specified version: %s is not installed", version)
	}

	// symlink to GoPath
	os.Remove(GoPath)
	if err := os.Symlink(fmt.Sprintf("%v/%v/go", GoInstallationDirectory, version), GoPath); err != nil {
		return err
	}

	// update DB
	versions := [][]byte{}
	if err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v not found", DBBucketName)
		}

		return bucket.ForEach(func(v, _ []byte) error {
			versions = append(versions, v)
			return nil
		})

	}); err != nil {
		return err
	}

	return DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(DBBucketName))
		if err != nil {
			return err
		}

		for _, v := range versions {
			value := ""
			if version == string(v) {
				value = DEFAULT
			}
			if err := bucket.Put(v, []byte(value)); err != nil {
				return err
			}
		}
		return nil
	})
}
