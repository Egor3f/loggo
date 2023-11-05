package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var BuildVersion string

// versionCmd represents the stream command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieves the build version",
	Long:  `Retrieves the build version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(BuildVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
