package bloc

import (
	"flean/services"
	"github.com/spf13/cobra"
)

// bloc/addCmd represents the bloc/add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a bloc inside the project",
	Long: `This command takes the bloc's name as argument and adds that bloc from the path specified in the flean.yaml file
The path from which to remove the bloc is stored inside the flean.yaml file and can be modified.
`,
	Example: "flean bloc add auth(auth is the name of the bloc)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blocName := args[0]
		services.GenerateBloc(blocName)
	},
}

func init() {
	Cmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bloc/addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bloc/addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
