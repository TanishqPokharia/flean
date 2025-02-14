/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package feature

import (
	"flean/utils"
	"github.com/spf13/cobra"
)

// Cmd represents the feature command
var Cmd = &cobra.Command{
	Use:   "feature",
	Short: "Manage features in a feature first architecture flutter app",
	Long: `This command is used to add or remove features from the features folder in a feature first architecture project.
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
	// feature/featureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feature/featureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
