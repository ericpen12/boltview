package cmd

import (
	"boltview/boltdb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			boltdb.Open(args[0])
		}
	},
}
