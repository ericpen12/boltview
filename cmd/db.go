package cmd

import (
	"boltview/boltdb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var dbCmd = &cobra.Command{
	Use:   "d",
	Short: "Use specific DB file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			boltdb.Open(args[0])
		}
	},
}
