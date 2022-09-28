package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

const (
	GO_DOWNLOAD_SERVER_URL = "https://go.dev/dl"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the go versions available",
	Long:  `Lists all the go version available`,
	Run: func(cmd *cobra.Command, args []string) {
		listGoVersions()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
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

func getAvailableGoVersions() goVersions {
	resp, err := http.Get(fmt.Sprintf("%v/?mode=json&include=all", GO_DOWNLOAD_SERVER_URL))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	gv := goVersions{}
	err = json.Unmarshal(body, &gv)
	if err != nil {
		panic(err)
	}
	return gv
}

func listGoVersions() {
	fmt.Printf("Listing Version for OS: %v ARCH: %v\n", termenv.String(runtime.GOOS).Italic(), termenv.String(runtime.GOARCH).Italic())

	for _, tag := range getAvailableGoVersions() {
		for _, pkg := range tag.Files {
			if pkg.Arch == runtime.GOARCH && pkg.Os == runtime.GOOS && pkg.Kind == "archive" {
				fmt.Println(pkg.Version)
			}
		}
	}
}

func getGoDownloadPackageInfo(version string) *packageInfo {
	for _, tag := range getAvailableGoVersions() {
		for _, pkg := range tag.Files {
			if pkg.Arch == runtime.GOARCH && pkg.Os == runtime.GOOS && pkg.Version == version && pkg.Kind == "archive" {
				return &pkg
			}
		}
	}
	return nil
}

func checkVersionInstalled(version string) bool {
	exists := false
	if err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v does not exist", version)
		}

		if bucket.Get([]byte(version)) != nil {
			exists = true
		}

		return nil
	}); err != nil {
		panic(err)
	}
	return exists
}

func checkVersionDefault(version string) bool {
	defaultVersion := false

	if err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(DBBucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %v does not exist", version)
		}

		if string(bucket.Get([]byte(version))) == "default" {
			defaultVersion = true
		}

		return nil
	}); err != nil {
		panic(err)
	}

	return defaultVersion
}
