package cmd

import (
	internalCommand "github.com/looCiprian/diff-backup/internal/command"
	"github.com/spf13/cobra"
)

var (

	initSource string

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init command will initialize the <source> directory that will used for backups",
		Long: `init command will initialize the <source> directory that will used for backups`,
		RunE: func(cmd *cobra.Command, args []string) error{

			internalCommand.SetInitConfig(initSource)
			if err := internalCommand.ExecuteInit(); err != nil {
				return err
			}

			return nil
		},
	}
)
