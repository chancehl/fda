package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// GetRunCommandFile determines the appropriate shell run command file (e.g., .bashrc or .zshrc).
// It returns the file path for the detected shell's RC file or defaults to .bashrc if the shell is unsupported.
// Returns an error if the user's home directory cannot be determined.
func GetRunCommandFile() (string, error) {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine user home directory: %w", err)
	}

	shell, err := GetShell()
	if err != nil {
		return "", err
	}

	switch shell {
	case "zsh":
		return path.Join(configDir, ".zshrc"), nil
	case "bash":
		return path.Join(configDir, ".bashrc"), nil
	default:
		fmt.Printf("Warning: Unsupported shell (%s), defaulting to .bashrc\n", shell)
		return path.Join(configDir, ".bashrc"), nil
	}
}

// GetShell retrieves the name of the user's shell by reading the SHELL environment variable.
// It returns only the executable name (e.g., "bash", "zsh") and an error if the shell cannot be determined.
func GetShell() (string, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return "", fmt.Errorf("could not determine shell")
	}

	// Extract just the executable name from the full path
	parts := strings.Split(shellPath, "/")
	return parts[len(parts)-1], nil
}
