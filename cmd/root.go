package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "diff-backup",
		Short: "diff-backup is a backup tool that perform incremental backup for a specific directory",
		Long: `diff-backup is a backup tool that perform incremental backup for a specific directory`,
	}
)

func init()  {

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&initSource, "source", "s","","Init directory")
	initCmd.MarkFlagRequired("source")

	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&backupSource, "source", "s", "", "Backup source")
	backupCmd.Flags().StringVarP(&backupDestination, "destination", "d", "", "Backup destination")
	backupCmd.MarkFlagRequired("source")
	backupCmd.MarkFlagRequired("destination")

	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVarP(&restoreSource, "source", "s", "", "Restore source")
	restoreCmd.Flags().StringVarP(&restoreDestination, "destination", "d", "", "Restore destination")
	restoreCmd.Flags().StringVarP(&restoreDate, "date", "t", "", "Restore date")
	restoreCmd.MarkFlagRequired("source")
	restoreCmd.MarkFlagRequired("destination")
	restoreCmd.MarkFlagRequired("date")

	restoreCmd.AddCommand(listRestoreDateCmd)
	listRestoreDateCmd.Flags().StringVarP(&restoreSource, "source", "s", "", "Restore source")
	listRestoreDateCmd.MarkFlagRequired("source")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		//fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
	}
}