package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Version is the current version of the Panda binary
const Version = "v0.0.1"

func init() {
	decorate("version", versionCmd)
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Panda",
	Long:  `All software has versions.  This is Panda's.`,
	Run: func(c *cobra.Command, args []string) {
		logrus.Infof("Panda version is: %s\n", Version)
	},
	// needed to override the root level pre-run func
	PersistentPreRunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}
