package main

import (
	"fmt"
	"os"
	"github.com/topherbullock/gifbase/cmd"
	"github.com/jessevdk/go-flags"
)

func main() {

	parser := flags.NewParser(&cmd.GifBaseCmd{}, flags.HelpFlag)
	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
