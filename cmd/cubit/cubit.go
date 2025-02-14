/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cubit

import (
	"flean/utils"
	"github.com/spf13/cobra"
)

// Cmd represents the cubit command
var Cmd = &cobra.Command{
	Use:   "cubit",
	Short: "Manage cubits inside your application",
	Long: `This command can be used with add & remove subcommands to manage cubits inside your application.
The path from which to add or remove the cubit is stored inside the flean.yaml file and can be modified.
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
	// cubit\cubitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cubit\cubitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
