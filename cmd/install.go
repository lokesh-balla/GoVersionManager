package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Specify the version of golang needed to install",
	Long: `Specify the version of golang which needs to be installed

	# To install a specific version
	$ gvm install go1.19
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			installGolang(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func progressBar(file string, size int, done chan bool) error {

	if size <= 0 {
		return fmt.Errorf("invalid size %v received", size)
	}

	// TODO: show a download bar
	// filename [====================>              ] 57%

	output := termenv.NewOutput(os.Stdout)
	output.SaveCursorPosition()

	t := time.NewTicker(time.Millisecond * 50)
	defer t.Stop()
	for {
		select {
		case <-done:
			output.ClearLine()
			output.RestoreCursorPosition()
			output.WriteString("100%\n")
			return nil
		case <-t.C:
			fileStat, err := os.Stat(file)
			if err != nil {
				return err
			}
			percent := float64(fileStat.Size()) / float64(size) * 100
			output.ClearLine()
			output.RestoreCursorPosition()
			output.WriteString(fmt.Sprintf("%.0f%%", percent))
		}
	}

}

func installGolang(version string) error {

	pInfo := getGoDownloadPackageInfo(version)
	if pInfo == nil {
		return fmt.Errorf("could not find go version: %v", version)
	}

	url := fmt.Sprintf("%v/%v", GO_DOWNLOAD_SERVER_URL, pInfo.Filename)

	headResp, err := http.Head(url)
	if err != nil {
		return err
	}
	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	done := make(chan bool)
	go func() {
		if err := progressBar(pInfo.Filename, size, done); err != nil {
			panic(err)
		}
	}()

	out, err := os.Create(pInfo.Filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	done <- true

	return nil
}
