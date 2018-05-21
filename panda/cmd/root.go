package cmd

import (
	"os"

	"github.com/ministryofprogramming/panda"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var anywhereCommands = []string{"version", "info", "help"}

var migrationPath string
var dialect string
var connectionString string

// RootCmd is the hook for all of the other Panda commands.
var RootCmd = &cobra.Command{
	SilenceErrors: true,
	Use:           "panda",
	Short:         "Robust schema evolution across all your environments.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		isFreeCommand := false
		for _, freeCmd := range anywhereCommands {
			if freeCmd == cmd.Name() {
				isFreeCommand = true
				continue
			}
		}

		if isFreeCommand {
			return nil
		}

		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main().
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&migrationPath, "path", "p", "./migrations", "Path to the migrations folder")
	RootCmd.PersistentFlags().StringVarP(&dialect, "dialect", "d", "mysql", "Database dialect, supported dialects are mysql, postgress, sqlite3.")
	RootCmd.PersistentFlags().StringVarP(&connectionString, "connection", "c", "", "Database connection string")
	RootCmd.PersistentFlags().BoolVarP(&panda.Debug, "verbose", "v", false, "Verbose logging")
	decorate("root", RootCmd)
}
