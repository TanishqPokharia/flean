package bloc

import (
	"flean/utils"
	"github.com/spf13/cobra"
)

var featureName *string

// Cmd represents the bloc command
var Cmd = &cobra.Command{
	Use:   "bloc",
	Short: "Manage blocs inside your application",
	Long: `This command can be used with add & remove subcommands to manage blocs inside your application.
The path from which to add or remove the bloc is stored inside the flean.yaml file and can be modified.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			utils.LogFatalError(err)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blocCmd.PersistentFlags().String("foo", "", "A help for foo")
	Cmd.PersistentFlags().StringP("feature", "f", "", "provide the feature name if using feature first architecture")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
