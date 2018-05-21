package migrate

import (
	"github.com/markbates/going/defaults"
	"github.com/ministryofprogramming/panda"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ResetCmd generates sql migration files
var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "The equivalent of running `migrate down` and then `migrate up`",
	RunE: func(cmd *cobra.Command, args []string) error {
		pFlag := cmd.Flag("path")
		migrationsPath := defaults.String(pFlag.Value.String(), "./migrations")

		cFlag := cmd.Flag("connection")
		connection := cFlag.Value.String()

		dFlag := cmd.Flag("dialect")
		dialect := dFlag.Value.String()

		conn, err := panda.Connect(dialect, connection)

		if err != nil {
			return errors.WithStack(err)
		}

		fm, err := panda.NewFileMigrator(migrationsPath, conn)
		if err != nil {
			return errors.WithStack(err)
		}
		return fm.Reset()
	},
}
