/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package feature

import (
	"flean/services"
	"flean/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a feature to your feature first clean architecture project inside the features folder",
	Long: `This command takes the feature's name as the argument and generates the data, domain and presentation layers
for that particular feature in the features folder.
`,
	Example: "flean feature add profile(profile is the name of the feature)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			utils.LogFatalError(err)
		}
		details, err := services.GetProjectDetails(path)
		path = fmt.Sprintf("%s/%s/lib/features/%s", details.Directory, details.Name, args[0])
		services.GenerateFFClean(path, args[0])
	},
}

func init() {
	Cmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feature/addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feature/addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
