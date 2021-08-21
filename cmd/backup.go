package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (

	backupSource string
	backupDestination string

	backupCmd = &cobra.Command{
		Use:   "backupCmd",
		Short: "backupCmd",
		Long: `backupCmd`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Source: " + backupSource + " destination: " + backupDestination)
		},
	}
)