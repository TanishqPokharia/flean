/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package feature

import (
	"flean/services"
	"flean/tui"
	"flean/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a feature to your ff clean architecture project inside the features folder",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			utils.LogFatalError(err)
		}
		details, err := services.GetProjectDetails(path)
		path = fmt.Sprintf("%s/%s/lib/features/%s", details.Directory, details.Name, args[0])
		services.GenerateFFClean(path)
		fmt.Println(tui.LogStyle.Render(fmt.Sprintf("%s feature added\n", args[0])))
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
