package cmd

import (
	internalCommand "github.com/looCiprian/diff-backup/internal/command"
	"fmt"
	"github.com/spf13/cobra"
)

var (

	restoreSource string
	restoreDestination string
	restoreDate string

	restoreCmd = &cobra.Command{
		Use:   "restore",
		Short: "restore command will restore the backup of the <source> directory to the <destination> directory from a certain <date>",
		Long: `restore command will restore the backup of the <source> directory to the <destination> directory from a certain <date>`,
		RunE: func(cmd *cobra.Command, args []string) error {

			internalCommand.SetRestoreConfig(restoreSource, restoreDestination, restoreDate)
			if err := internalCommand.ExecuteRestore(); err != nil {
				return err
			}

			return nil
		},
	}

	listRestoreDateCmd = &cobra.Command{
		Use:   "listDates",
		Short: "listDates command will list the available dates that can be restores",
		Long: `listDates command will list the available dates that can be restores`,
		RunE: func(cmd *cobra.Command, args []string) error {

			internalCommand.SetRestoreConfig(restoreSource, "", "")
			dates , err := internalCommand.GetResorableDates()

			if err != nil {
				return err
			}

			for _, date := range dates {
				fmt.Println("- " + date)
			}

			return nil
		},
	}
)