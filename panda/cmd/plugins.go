package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"github.com/ministryofprogramming/panda/env"
	"github.com/ministryofprogramming/panda/plugins"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var plugx = &sync.Mutex{}
var _plugs plugins.List

func plugs() plugins.List {
	plugx.Lock()
	defer plugx.Unlock()
	if _plugs == nil {
		var err error
		_plugs, err = plugins.Available()
		if err != nil {
			_plugs = plugins.List{}
			logrus.Errorf("error loading plugins %s\n", err)
		}
	}
	return _plugs
}

func decorate(name string, cmd *cobra.Command) {
	pugs := plugs()
	for _, c := range pugs[name] {
		func(c plugins.Command) {
			anywhereCommands = append(anywhereCommands, c.Name)
			cc := &cobra.Command{
				Use:     c.Name,
				Short:   fmt.Sprintf("[PLUGIN] %s", c.Description),
				Aliases: c.Aliases,
				RunE: func(cmd *cobra.Command, args []string) error {
					plugCmd := c.Name
					if c.UseCommand != "" {
						plugCmd = c.UseCommand
					}

					ax := []string{plugCmd}
					if plugCmd == "-" {
						ax = []string{}
					}

					ax = append(ax, args...)
					ex := exec.Command(c.Binary, ax...)
					if runtime.GOOS != "windows" {
						ex.Env = append(env.Environ(), "COMMANDER_PLUGIN=1")
					}
					ex.Stdin = os.Stdin
					ex.Stdout = os.Stdout
					ex.Stderr = os.Stderr
					return ex.Run()
				},
			}
			cc.DisableFlagParsing = true
			cmd.AddCommand(cc)
		}(c)
	}
}
