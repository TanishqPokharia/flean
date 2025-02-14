/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cubit

import (
	"flean/services"
	"github.com/spf13/cobra"
)

// addCmd represents the cubit\add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a cubit inside the project",
	Long: `This command takes the cubit's name as argument and adds that cubit from the path specified in the flean.yaml file
The path from which to add the cubit is stored inside the flean.yaml file and can be modified.
`,
	Example: "flean cubit add auth (auth is the name of the cubit)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		services.GenerateCubit(args[0])
	},
}

func init() {
	Cmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cubit\addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cubit\addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
