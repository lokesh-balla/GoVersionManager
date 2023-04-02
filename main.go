package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Lokesh-Balla/GoVersionManager/cmd"
)

var (
	version string
	build   string
)

func main() {

	if version == "" {
		version = "dev"
	}
	if build == "" {
		build = time.Now().Format("20060102")
	}

	if err := cmd.Execute(version, build); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cmd.DB.Close()
}
