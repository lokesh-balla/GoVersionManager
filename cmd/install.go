package cmd

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	TerminalWidth  int
	TerminalHeight int
)

func init() {
	TerminalWidth, TerminalHeight, _ = term.GetSize(int(os.Stdout.Fd()))
}

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
			if err := installGolang(args[0]); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func drawProgressBar(output *termenv.Output, file string, width int, percent float64) {
	if width < 0 || width > 80 {
		width = 80
	}
	drawString := fmt.Sprintf("%s [", file)
	width = width - len(drawString) - 10
	barPercent := (percent / 100) * float64(width)
	for i := 0; i < width; i++ {
		if i <= int(barPercent) {
			drawString = fmt.Sprintf("%s%s", drawString, "=")
		} else {
			drawString = fmt.Sprintf("%s%s", drawString, " ")
		}
	}
	drawString = fmt.Sprintf("%s] %.0f%%", drawString, percent)
	output.WriteString(drawString)
}

func progressBar(file string, size int, done chan bool) error {

	if size <= 0 {
		return fmt.Errorf("invalid size %v received", size)
	}

	output := termenv.NewOutput(os.Stdout)
	output.SaveCursorPosition()

	t := time.NewTicker(time.Millisecond * 50)
	defer t.Stop()
	for {
		select {
		case <-done:
			output.ClearLine()
			output.RestoreCursorPosition()
			drawProgressBar(output, file, TerminalWidth, 100)
			return nil
		case <-t.C:
			fileStat, err := os.Stat(file)
			if err != nil {
				return err
			}
			percent := float64(fileStat.Size()) / float64(size) * 100
			output.ClearLineLeft()
			output.RestoreCursorPosition()
			drawProgressBar(output, file, TerminalWidth, percent)
		}
	}

}

func downloadGolang(filePath string, url string) error {
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
		if err := progressBar(filePath, size, done); err != nil {
			panic(err)
		}
	}()

	out, err := os.Create(filePath)
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

func installGolang(version string) error {

	pInfo := getGoDownloadPackageInfo(version)
	if pInfo == nil {
		return fmt.Errorf("could not find go version: %v", version)
	}

	url := fmt.Sprintf("%v/%v", GO_DOWNLOAD_SERVER_URL, pInfo.Filename)

	filePath := fmt.Sprintf("%s/%s", GoInstallationDirectory, pInfo.Filename)

	if _, err := os.Stat(filePath); err != nil {
		if err := downloadGolang(filePath, url); err != nil {
			return err
		}
	}

	// check sha256 sum
	valid, err := checkSHA256Sum(filePath, pInfo.Sha256)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("downloaded file has invalid sha256sum")
	}

	// extract Go
	if err := extractGoTar(filePath, fmt.Sprintf("%s/%s", GoInstallationDirectory, version)); err != nil {
		return err
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

	return nil
}

func extractGoTar(filePath string, targetPath string) error {

	madeDir := map[string]bool{}

	reader, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("not a gzip compressed body: %v", err)
	}

	tarReader := tar.NewReader(gzipReader)

	for {
		f, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar error: %v", err)
		}
		rel := filepath.FromSlash(f.Name)
		abs := filepath.Join(targetPath, rel)

		fi := f.FileInfo()
		mode := fi.Mode()
		switch {
		case mode.IsRegular():
			dir := filepath.Dir(abs)
			if !madeDir[dir] {
				if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
					return err
				}
				madeDir[dir] = true
			}
			wf, err := os.OpenFile(abs, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode.Perm())
			if err != nil {
				return err
			}
			n, err := io.Copy(wf, tarReader)
			if closeErr := wf.Close(); closeErr != nil && err == nil {
				err = closeErr
			}
			if err != nil {
				return fmt.Errorf("error writing to %s: %v", abs, err)
			}
			if n != f.Size {
				return fmt.Errorf("only wrote %d bytes to %s; expected %d", n, abs, f.Size)
			}
		case mode.IsDir():
			if err := os.MkdirAll(abs, 0755); err != nil {
				return err
			}
			madeDir[abs] = true
		default:
			return fmt.Errorf("tar file entry %s contained unsupported file type %v", f.Name, mode)
		}
	}

	return nil
}

func checkSHA256Sum(filePath string, SHA256 string) (bool, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return false, err
	}

	if SHA256 != hex.EncodeToString(hash.Sum(nil)) {
		return false, nil
	}

	return true, nil
}
