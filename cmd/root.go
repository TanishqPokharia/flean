package cmd

import (
	"flean/cmd/bloc"
	"flean/cmd/cubit"
	"flean/cmd/feature"
	"flean/constants"
	"flean/tui"
	"flean/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flean",
	Short: fmt.Sprintf("%s\nFlean - Open-source Flutter app architecture generator and manager\n", tui.LogStyle.Bold(true).Render(constants.Logo)),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			utils.LogFatalError(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bloc.Cmd)
	rootCmd.AddCommand(cubit.Cmd)
	rootCmd.AddCommand(feature.Cmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flean.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
