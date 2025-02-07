package cubit

import (
	"flean/services"
	"github.com/spf13/cobra"
)

// cubit\removeCmd represents the cubit\remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a cubit from the application",
	Long: `
	This command takes the cubit's name as argument and removes that cubit from the path specified in the flean.yaml file.
	The path from which to remove the cubit is stored inside the flean.yaml file and can be modified.
	Sample Usage:
		flean cubit remove auth
'
`,
	Run: func(cmd *cobra.Command, args []string) {
		services.RemoveCubit(args[0])
	},
}

func init() {
	Cmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cubit\removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cubit\removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
