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
	Short: "Removes the entire feature from the features folder inside feature first architecture project",
	Long: `This command takes the feature's name as the argument and removes that feature including it's data, domain and presentation layers
for that particular feature from the features folder.
`,
	Example: "flean feature remove profile(profile is the name of the feature)",
	Args:    cobra.ExactArgs(1),
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
