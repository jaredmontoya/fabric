package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/danielmiessler/fabric/pkg/cli"
)

func main() {
	err := cli.Cli(version)
	if err != nil && !flags.WroteHelp(err) {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
