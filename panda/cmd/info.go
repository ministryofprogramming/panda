package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ministryofprogramming/panda/env"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints off diagnostic information useful for debugging.",
	RunE: func(cmd *cobra.Command, args []string) error {
		bb := os.Stdout

		bb.WriteString(fmt.Sprintf("### Panda Version\n%s\n", Version))

		return runInfoCmds()
	},
}

type infoCommand struct {
	Name      string
	PathName  string
	Cmd       *exec.Cmd
	InfoLabel string
}

func runInfoCmds() error {

	commands := []infoCommand{
		{"Go", env.Get("GO_BIN", "go"), exec.Command(env.Get("GO_BIN", "go"), "version"), "\n### Go Version\n"},
		{"Go", env.Get("GO_BIN", "go"), exec.Command(env.Get("GO_BIN", "go"), "env"), "\n### Go Env\n"},
		{"dep", "dep", exec.Command("dep", "version"), "\n### Dep Version\n"},
		{"dep", "dep", exec.Command("dep", "status"), "\n### Dep Status\n"},
	}

	for _, cmd := range commands {
		err := execIfExists(cmd)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func execIfExists(infoCmd infoCommand) error {
	bb := os.Stdout
	bb.WriteString(infoCmd.InfoLabel)

	if _, err := exec.LookPath(infoCmd.PathName); err != nil {
		bb.WriteString(fmt.Sprintf("%s Not Found\n", infoCmd.Name))
		return nil
	}

	infoCmd.Cmd.Stdout = bb
	infoCmd.Cmd.Stderr = bb

	err := infoCmd.Cmd.Run()
	return err
}

func init() {
	decorate("info", RootCmd)
	RootCmd.AddCommand(infoCmd)
}
