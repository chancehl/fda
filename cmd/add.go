package cmd

import (
	"fmt"
	"os"

	"github.com/chancehl/fda/internal/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:  "add [name]",
	Args: cobra.ExactArgs(1),
	RunE: execute,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func execute(cmd *cobra.Command, args []string) error {
	name := args[0]

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current directory and dir parameter was not specified: %w", err)
	}

	alias, err := models.NewAlias(name, wd, runCommandFile)
	if err != nil {
		return err
	}

	if err := alias.WriteToFile(); err != nil {
		return err
	}

	fmt.Printf("[OK] added go-%s alias to %s\n", alias.Name, runCommandFile)
	return nil
}
