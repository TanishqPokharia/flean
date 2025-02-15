package bloc

import (
	"flean/services"
	"flean/utils"
	"github.com/spf13/cobra"
)

// bloc/removeCmd represents the bloc/remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an entire bloc from the project",
	Long: `This command takes the bloc's name as argument and removes that bloc from the path specified in the flean.yaml file.
The path from which to remove the bloc is stored inside the flean.yaml file and can be modified.
`,
	Example: "flean bloc remove auth(auth is the name of the bloc)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		featureName, err := cmd.Flags().GetString("feature")
		if err != nil {
			utils.LogFatalError(err)
		}
		services.RemoveBloc(args[0], featureName)
	},
}

func init() {
	Cmd.AddCommand(removeCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bloc/removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bloc/removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
