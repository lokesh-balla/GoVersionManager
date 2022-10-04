package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sort"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

const (
	GO_DOWNLOAD_SERVER_URL = "https://go.dev/dl"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:           "list",
	Short:         "Lists all the go versions installed",
	Long:          `Lists all the go version installed`,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		getAll, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}
		err = listGoVersions(getAll)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "list all the available versions")
}

type goVersions []struct {
	Version string        `json:"version"`
	Stable  bool          `json:"stable"`
	Files   []packageInfo `json:"files"`
}

type packageInfo struct {
	Filename string `json:"filename"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
	Sha256   string `json:"sha256"`
	Size     int    `json:"size"`
	Kind     string `json:"kind"`
}

func getAvailableGoVersions() (goVersions, error) {
	resp, err := http.Get(fmt.Sprintf("%v/?mode=json&include=all", GO_DOWNLOAD_SERVER_URL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	gv := goVersions{}
	err = json.Unmarshal(body, &gv)
	if err != nil {
		return nil, err
	}
	return gv, nil
}

func listGoVersions(all bool) error {

	fmt.Printf("OS: %v ARCH: %v\n\n", termenv.String(runtime.GOOS).Italic().Foreground(termenv.ANSIGreen), termenv.String(runtime.GOARCH).Italic().Foreground(termenv.ANSIGreen))

	fmt.Println(termenv.String("Installed Versions").Bold().Foreground(termenv.ANSIBlue))

	goVersions, err := listAllInstalledVersions()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		goVersions = map[string]string{}
	}

	versions := []string{}
	for v := range goVersions {
		versions = append(versions, v)
	}

	sort.Slice(versions, func(i, j int) bool { return versions[i] > versions[j] })
	for _, v := range versions {
		if goVersions[v] == DEFAULT {
			fmt.Println(termenv.String(v, "← default").Bold())
			continue
		}
		fmt.Println(termenv.String(v).Faint())
	}
	versions = []string{}
	if all {
		availableGoVersions, err := getAvailableGoVersions()
		if err != nil {
			return err
		}
		for _, tag := range availableGoVersions {
			for _, pkg := range tag.Files {
				if pkg.Arch == runtime.GOARCH && pkg.Os == runtime.GOOS && pkg.Kind == "archive" {
					if _, ok := goVersions[pkg.Version]; !ok {
						versions = append(versions, pkg.Version)
					}
				}
			}
		}

		sort.Slice(versions, func(i, j int) bool { return versions[i] > versions[j] })

		fmt.Println(termenv.String("Available Versions").Bold().Foreground(termenv.ANSIBlue))
		for _, v := range versions {
			fmt.Println(termenv.String(v).Faint())
		}
	}

	return nil
}

func getGoDownloadPackageInfo(version string) (*packageInfo, error) {
	availableGoVersions, err := getAvailableGoVersions()
	if err != nil {
		return nil, err
	}

	for _, tag := range availableGoVersions {
		for _, pkg := range tag.Files {
			if pkg.Arch == runtime.GOARCH && pkg.Os == runtime.GOOS && pkg.Version == version && pkg.Kind == "archive" {
				return &pkg, nil
			}
		}
	}
	return nil, fmt.Errorf("could not find go version: %v", version)
}

func listAllInstalledVersions() (map[string]string, error) {
	m := map[string]string{}
	if err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v does not exist", DBBucketName)
		}

		bucket.ForEach(func(key, value []byte) error {
			m[string(key)] = string(value)
			return nil
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return m, nil
}

func checkVersionInstalled(version string) (bool, error) {
	exists := false
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v does not exist", DBBucketName)
		}

		if bucket.Get([]byte(version)) != nil {
			exists = true
		}

		return nil
	})
	return exists, err
}

func checkVersionDefault(version string) (bool, error) {
	defaultVersion := false

	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v does not exist", DBBucketName)
		}

		if string(bucket.Get([]byte(version))) == DEFAULT {
			defaultVersion = true
		}

		return nil
	})

	return defaultVersion, err
}
