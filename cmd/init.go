package cmd

import (
	internalCommand "diff-backup/internal/command"
	"fmt"
	"github.com/spf13/cobra"
)

var (

	initSource string

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init command will initialize the <source> directory that will used for backups",
		Long: `init command will initialize the <source> directory that will used for backups`,
		Run: func(cmd *cobra.Command, args []string) {

			internalCommand.SetInitConfig(initSource)
			if err := internalCommand.ExecuteInit(); err != nil {
				fmt.Println(err)
			}

		},
	}
)
