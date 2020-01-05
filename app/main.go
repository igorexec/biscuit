package main

import (
	"fmt"
	"github.com/igorexec/biscuit/app/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

type Opts struct {
	ServerCmd cmd.ServerCommand `command:"server"`
}

var revision = "unknown"

func main() {
	fmt.Printf("Biscuit version: %s", revision)

	var opts Opts
	p := flags.NewParser(&opts, flags.Default)

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.HelpFlag {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
