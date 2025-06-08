package main

import (
	"github.com/black-dev-x/cobra-cli/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.Execute()
}
