package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "[error]", err)
		os.Exit(1)
	}
}

func run() error {
	args, err := ParseArgs()
	if err != nil {
		return err
	}

	runCommandFile, err := GetRunCommandFile()
	if err != nil {
		return err
	}

	alias, err := NewAlias(args.Name, args.Dir, runCommandFile)
	if err != nil {
		return err
	}

	if err := alias.WriteToFile(); err != nil {
		return err
	}

	fmt.Printf("[success] Added go-%s alias to %s\n", alias.Name, runCommandFile)
	return nil
}
