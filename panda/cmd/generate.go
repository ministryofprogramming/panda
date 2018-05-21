package cmd

import (
	"github.com/ministryofprogramming/panda/panda/cmd/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generates database related scripts.",
}

func init() {
	generateCmd.AddCommand(generate.MigrationCmd)

	RootCmd.AddCommand(generateCmd)
}
