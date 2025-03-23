package main

import (
	"fmt"
	"os"

	"sitegenerator/cli"
)

func main() {
	rootCmd := cli.CreateRootCommand(AppVersion)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
