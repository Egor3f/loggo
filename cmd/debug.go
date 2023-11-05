package cmd

import (
	"github.com/aurc/loggo/internal/loggo"
	"github.com/aurc/loggo/internal/reader"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Continuously stream l'oggo log",
	Long: `This command aims to assist troubleshoot loggos issue and would be rarely utilised by loggo's users':

	loggo debug`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := reader.MakeReader(loggo.LatestLog, nil)
		app := loggo.NewLoggoApp(reader, "")
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
