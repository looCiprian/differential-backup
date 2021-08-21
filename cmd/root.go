package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (

	rootCmd = &cobra.Command{
		Use:   "diff-backup",
		Short: "diff-backup",
		Long: `diff-backup`,
	}
)

func init()  {
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

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}