package cmd

import (
	internalCommand "github.com/looCiprian/diff-backup/internal/command"
	"github.com/spf13/cobra"
)

var (

	backupSource string
	backupDestination string

	backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "backup command will perform the backup of the <source> directory to the <destination> directory",
		Long: `backup command will perform the backup of the <source> directory to the <destination> directory`,
		RunE: func(cmd *cobra.Command, args []string) error {

			internalCommand.SetBackupConfig(backupSource, backupDestination)
			if err := internalCommand.ExecuteBackup(); err != nil {
				return err
			}

			return nil
		},
	}
)