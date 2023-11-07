package cmd

import (
	"github.com/egor3f/loggo/internal/reader"
	"os"

	"github.com/egor3f/loggo/internal/loggo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loggo",
	Short: "Continuously stream log input source",
	Long: `Continuously stream log entries from an input stream such
as the standard input (through pipe) or a input file. Note that
if it's reading from a file, it automatically detects file 
rotation and continue to stream. For example:

	loggo --file <file-path>
	<some arbitrary input> | loggo`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := cmd.Flag("file").Value.String()
		templateFile := cmd.Flag("template").Value.String()
		reader := reader.MakeReader(fileName, nil)
		app := loggo.NewLoggoApp(reader, templateFile)
		app.Run()
	},
}

// Initiate adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Initiate() {
	loggo.BuildVersion = BuildVersion
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().
		StringP("file", "f", "", "Input Log File")
	rootCmd.Flags().
		StringP("template", "t", "", "Rendering Template")
}
