package cmd

import (
	internalCommand "diff-backup/internal/command"
	"fmt"
	"github.com/spf13/cobra"
)

var (

	backupSource string
	backupDestination string

	backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "backup command will perform the backup of the <source> directory to the <destination> directory",
		Long: `backup command will perform the backup of the <source> directory to the <destination> directory`,
		Run: func(cmd *cobra.Command, args []string) {
			internalCommand.SetBackupConfig(backupSource, backupDestination)
			if err := internalCommand.ExecuteBackup(); err != nil {
				fmt.Println(err)
			}

		},
	}
)