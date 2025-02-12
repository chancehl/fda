package main

import (
	"fmt"
	"os"

	"github.com/chancehl/fda/internal/alias"
	"github.com/chancehl/fda/internal/arguments"
	"github.com/chancehl/fda/internal/environment"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "[error]", err)
		os.Exit(1)
	}
}

func run() error {
	args, err := arguments.Parse()
	if err != nil {
		return err
	}

	runCommandFile, err := environment.GetRunCommandFile()
	if err != nil {
		return err
	}

	alias, err := alias.New(args.Name, args.Dir, runCommandFile)
	if err != nil {
		return err
	}

	if err := alias.WriteToFile(); err != nil {
		return err
	}

	fmt.Printf("[success] Added go-%s alias to %s\n", alias.Name, runCommandFile)
	return nil
}
