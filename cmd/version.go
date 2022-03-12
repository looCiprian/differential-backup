package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of diff-backup",
	Long:  `All software has versions. This is diff-backup's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diff-backup v3.1 -- HEAD")
	},
}
