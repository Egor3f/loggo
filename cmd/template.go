package cmd

import (
	"github.com/egor3f/loggo/internal/config"
	"github.com/egor3f/loggo/internal/loggo"
	"github.com/egor3f/loggo/internal/util"
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Starts the loggo template manager app only",
	Long: `Starts the loggo template manager app only so that
you can edit or create new templates. For example:

To start from a blank canvas:
	loggo template
To start from an existing file and update or save new from it:
	loggo template --file <some existing template>
To start from an example template:
	loggo template --example=true
`,
	Run: func(cmd *cobra.Command, args []string) {
		templateFile := cmd.Flag("file").Value.String()
		example := cmd.Flag("example").Value.String() == "true"
		var cfg *config.Config
		var err error
		if len(templateFile) == 0 {
			if example {
				cfg, err = config.MakeConfig("")
			} else {
				cfg = &config.Config{
					Keys:          make([]config.Key, 0),
					LastSavedName: "",
				}
			}
		} else {
			cfg, err = config.MakeConfig(templateFile)
		}
		if err != nil {
			util.Log().Fatal("Unable to start app: ", err)
		}
		app := loggo.NewAppWithConfig(cfg)
		view := loggo.NewTemplateView(app, true, nil, nil)
		app.Run(view)

	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().
		StringP("file", "f", "", "Input Template File")
	templateCmd.Flags().
		StringP("example", "e", "", "Load example log template. "+
			"If `file` flag provided this flag is ignored.")
}
