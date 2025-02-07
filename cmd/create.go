package cmd

import (
	"flean/constants"
	"flean/services"
	"flean/tui"
	"flean/tui/models"
	"flean/utils"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"slices"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new flutter project with desired architecture",
	Long:    ``,
	Args:    cobra.ExactArgs(1),
	Example: "  flean create my_flutter_app",
	Run: func(cmd *cobra.Command, args []string) {
		startTui(args[0])
	},
}

func startTui(name string) {
	options := []string{"Clean", "Feature First Clean", "MVVM", "MVC"}
	var selectedOption string
	arch := tea.NewProgram(models.OptionsModel{
		Question: "Select Architecture:",
		Options:  options,
		OnSelect: func(index int) {
			selectedOption = options[index]
			fmt.Println(tui.LogStyle.Render(fmt.Sprintf("Creating in %s Architecture\n", options[index])))
		},
	})
	if _, err := arch.Run(); err != nil {
		utils.LogFatalError(err)
	}
	var dependencies []string
	dep := tea.NewProgram(models.MultiOptionsModel{
		Question: "Select commonly used packages to include",
		Options:  constants.Dependencies,
		Headers:  []string{"No.", "Package", "Category"},
		OnSelect: func(index int, option string) []string {
			optIndex := slices.Index(dependencies, option)
			if optIndex != -1 {
				dependencies = append(dependencies[:optIndex], dependencies[optIndex+1:]...)
			} else {
				dependencies = append(dependencies, option)
			}
			return dependencies
		},
	}, tea.WithMouseAllMotion(), tea.WithMouseCellMotion())
	if _, err := dep.Run(); err != nil {
		utils.LogFatalError(err)
	}
	loader := tea.NewProgram(models.NewLoader("Adding dependencies...", func() (msg tea.Msg) {
		err := services.GenerateBaseApp(name)
		if err != nil {
			fmt.Println(err)
			return tea.QuitMsg{}
		}
		switch selectedOption {
		case "Clean":
			services.GenerateClean(name)
		case "Feature First Clean":
			services.GenerateFFBase(name)
		case "MVVM":
			services.GenerateMVVM(name)
		case "MVC":
			services.GenerateMVC(name)
		}
		err = services.AddDependencies(name, dependencies...)
		if err != nil {
			fmt.Println(err)
			return tea.QuitMsg{}
		}
		err = services.StoreProjectDetails(name, selectedOption)
		if err != nil {
			fmt.Println(err)
			return tea.QuitMsg{}
		}
		if err != nil {
			fmt.Println(err)
			return tea.QuitMsg{}
		}
		fmt.Println(tui.LogStyle.Render(fmt.Sprintf("Build Success!\n%s created in %s Architecture\n", name, selectedOption)))
		return tea.QuitMsg{}
	}))
	if _, err := loader.Run(); err != nil {
		utils.LogFatalError(err)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
