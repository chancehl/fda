package cmd

import (
	"os"

	"github.com/chancehl/fda/internal"
	"github.com/spf13/cobra"
)

var runCommandFile string

var rootCmd = &cobra.Command{
	Use:              "fda",
	Short:            "fda creates directory aliases for you to use",
	PersistentPreRun: PersistentPreRun,
}

func PersistentPreRun(cmd *cobra.Command, args []string) {
	if file, err := internal.GetRunCommandFile(); err == nil {
		runCommandFile = file
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
