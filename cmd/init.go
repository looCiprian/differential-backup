package cmd

import (
	"github.com/spf13/cobra"
)

var (

	initSource string

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init",
		Long: `init`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)
