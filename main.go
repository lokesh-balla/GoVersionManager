package main

import (
	"github.com/Lokesh-Balla/GoVersionManager/cmd"
)

func main() {
	cmd.Execute()
	defer cmd.DB.Close()
}
