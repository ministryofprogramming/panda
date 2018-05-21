package migrate

import (
	"github.com/ministryofprogramming/panda"
	"github.com/pkg/errors"

	"github.com/markbates/going/defaults"
	// we import all dialects as we dont know at this point which will be used
	_ "github.com/ministryofprogramming/panda/dialects/mysql"
	_ "github.com/ministryofprogramming/panda/dialects/postgress"
	_ "github.com/ministryofprogramming/panda/dialects/sqlite3"
	"github.com/spf13/cobra"
)

// UpCmd generates sql migration files
var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all of the 'up' migrations.",
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
		return fm.Up()
	},
}
