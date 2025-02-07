/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package feature

import (
	"flean/services"
	"flean/tui"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := services.RemoveFeatureClean(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tui.LogStyle.Render(fmt.Sprintf("%s feature removed\n", args[0])))
	},
}

func init() {
	Cmd.AddCommand(removeCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feature/removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feature/removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
