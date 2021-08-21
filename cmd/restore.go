package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (

	restoreSource string
	restoreDestination string
	restoreDate string

	restoreCmd = &cobra.Command{
		Use:   "restoreCmd",
		Short: "restoreCmd",
		Long: `restoreCmd`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Source: " + restoreSource + " destination: " + restoreDestination)
		},
	}

	listRestoreDateCmd = &cobra.Command{
		Use:   "listrestoreCmd",
		Short: "listrestoreCmd",
		Long: `listrestoreCmd`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)