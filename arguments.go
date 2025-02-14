package main

import (
	"errors"
	"fmt"
	"os"
)

// Args represents the parsed command-line arguments.
type Args struct {
	// Name is the alias name provided by the user.
	Name string
	// Dir is the directory associated with the alias.
	Dir string
}

// ParseArgs extracts command-line arguments and returns an Args struct.
// It expects at least one argument (the alias name).
// If a second argument (directory) is not provided, it defaults to the current working directory.
// Returns an error if the alias name is missing or if the specified directory does not exist.
func ParseArgs() (*Args, error) {
	var name string
	if len(os.Args) <= 1 {
		return nil, fmt.Errorf("missing required parameter <name>")
	} else {
		name = os.Args[1]
	}

	var dir string
	if len(os.Args) <= 2 {
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("could not get current directory and dir parameter was not specified: %w", err)
		}
		dir = wd
	} else {
		if _, err := os.Stat(os.Args[2]); errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("directory %s does not exist", os.Args[2])
		} else {
			dir = os.Args[2]
		}
	}

	return &Args{Name: name, Dir: dir}, nil
}
