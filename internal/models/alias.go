package models

import (
	"fmt"
	"os"
	"regexp"
)

// Alias represents a shell alias that maps a user-friendly name to a command.
type Alias struct {
	// Name is the user-defined alias name.
	Name string
	// Command is the actual shell command that will be assigned to the alias.
	Command string
	// RunCommandFile is the shell profile file where the alias will be written (e.g., .bashrc, .zshrc).
	RunCommandFile string
}

// NewAlias creates a new Alias instance after validating the alias name and checking for duplicates in the shell profile.
// - `name` is the desired alias name.
// - `dir` is the directory the alias should navigate to.
// - `runCommandFile` is the shell profile file where the alias should be stored.
// Returns an error if the alias name is invalid, already exists, or if the shell profile cannot be read.
func NewAlias(name, dir, runCommandFile string) (*Alias, error) {
	// validate alias name (allow only alphanumeric characters, underscores, and hyphens)
	matched, err := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, name)
	if err != nil {
		return nil, fmt.Errorf("could not validate alias name: %w", err)
	}
	if !matched {
		return nil, fmt.Errorf("invalid alias name: only letters, numbers, underscores, and hyphens are allowed")
	}

	// check if the alias already exists in the shell profile
	bytes, err := os.ReadFile(runCommandFile)
	if err != nil {
		return nil, fmt.Errorf("could not read shell profile to check if alias already exists: %w", err)
	}

	re, err := regexp.Compile(fmt.Sprintf(`(?m)^\s*alias go-%s=.*`, name))
	if err != nil {
		return nil, fmt.Errorf("could not compile regex to check for alias validity: %w", err)
	}
	if re.Match(bytes) {
		return nil, fmt.Errorf("alias go-%s already exists in run command file", name)
	}

	return &Alias{
		Name:           name,
		RunCommandFile: runCommandFile,
		Command:        fmt.Sprintf(`alias go-%s="cd %s"`, name, dir),
	}, nil
}

// WriteToFile appends the alias command to the shell profile file.
// Returns an error if the file cannot be opened or written to.
func (a *Alias) WriteToFile() error {
	flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE

	f, err := os.OpenFile(a.RunCommandFile, flags, 0600)
	if err != nil {
		return fmt.Errorf("failed to open shell profile for writing: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(a.Command + "\n"); err != nil {
		return fmt.Errorf("failed to write alias to profile: %w", err)
	}
	return nil
}
