package cmd

import (
	"github.com/aurc/loggo/internal/loggo"
	"github.com/aurc/loggo/internal/reader"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Continuously stream log input source",
	Long: `Continuously stream log entries from an input stream such
as the standard input (through pipe) or a input file. Note that
if it's reading from a file, it automatically detects file 
rotation and continue to stream. For example:

	loggo stream --file <file-path>
	<some arbitrary input> | loggo stream`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := cmd.Flag("file").Value.String()
		templateFile := cmd.Flag("template").Value.String()
		reader := reader.MakeReader(fileName, nil)
		app := loggo.NewLoggoApp(reader, templateFile)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
	streamCmd.Flags().
		StringP("file", "f", "", "Input Log File")
	streamCmd.Flags().
		StringP("template", "t", "", "Rendering Template")
}
