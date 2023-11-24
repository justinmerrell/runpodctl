package project

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectName string
var modelType string
var modelName string
var initCurrentDir bool

const inputPromptPrefix string = "   > "

func prompt(message string) string {
	var s string = ""
	for s == "" {
		fmt.Print(inputPromptPrefix + message + ": ")
		fmt.Scanln(&s)
	}
	return s
}
func contains(input string, choices []string) bool {
	for _, choice := range choices {
		if input == choice {
			return true
		}
	}
	return false
}
func promptChoice(message string, choices []string, defaultChoice string) string {
	var s string = ""
	for !contains(s, choices) {
		fmt.Print(inputPromptPrefix + message + " (" + strings.Join(choices, ", ") + ") " + "[" + defaultChoice + "]" + ": ")
		fmt.Scanln(&s)
		if s == "" {
			return defaultChoice
		}
	}
	return s
}

var NewProjectCmd = &cobra.Command{
	Use:   "new",
	Args:  cobra.ExactArgs(0),
	Short: "create a new project",
	Long:  "create a new Runpod project folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a new project...")
		//prompt project name (if not provided)
		if projectName == "" {
			projectName = prompt("Enter the project name")
		} else {
			fmt.Println("Project name: " + projectName)
		}
		promptTemplates := &promptui.SelectTemplates{
			Help:     "",
			Label:    inputPromptPrefix + "{{ . }}",
			Active:   ` ● {{ . | cyan }}`,
			Inactive: ` ○ {{ . | white }}`,
			Selected: `{{ "✔" | green }} {{ . | white }}`,
		}
		getNetworkVolume := promptui.Select{
			Label:     "Select a Network Volume",
			Items:     []string{"Option1", "Option2", "Option3"},
			Templates: promptTemplates,
		}
		_, networkVolumeId, err := getNetworkVolume.Run()
		if err != nil {
			//ctrl c for example
			return
		}
		cudaVersion := promptChoice("Select a CUDA version, or press enter to use the default",
			[]string{"11.1.1", "11.8.0", "12.1.0"}, "11.8.0")
		pythonVersion := promptChoice("Select a Python version, or press enter to use the default",
			[]string{"3.8", "3.9", "3.10", "3.11"}, "3.10")
		fmt.Println(networkVolumeId, cudaVersion, pythonVersion)

		fmt.Printf(`
Project Summary:
   - Project Name: %s
   - RunPod Network Storage ID: %s
   - CUDA Version: %s
   - Python Version: %s
		`, projectName, networkVolumeId, cudaVersion, pythonVersion)
		fmt.Println()
		//create files
		//folder structure (check for --init)
		//project toml
	},
}

var StartProjectCmd = &cobra.Command{
	Use:   "start",
	Args:  cobra.ExactArgs(0),
	Short: "start current project",
	Long:  "start a development pod session for the Runpod project in the current folder",
	Run: func(cmd *cobra.Command, args []string) {
		//parse project toml
		//check for existing pod or
		//try to get pod with one of gpu types
		//open ssh connection
		//create remote folder structure
		//rsync project files
		//activate venv on remote
		//create file watcher
		//run launch api server / hot reload loop
	},
}

var DeployProjectCmd = &cobra.Command{
	Use:   "deploy",
	Args:  cobra.ExactArgs(0),
	Short: "deploy current project",
	Long:  "deploy an endpoint for the Runpod project in the current folder",
	Run: func(cmd *cobra.Command, args []string) {
		//parse project toml
		//check for existing pod or
		//try to get pod with one of gpu types
		//open ssh connection
		//sync remote dev to remote prod
		//deploy new template
		//deploy / update endpoint
	},
}

// var BuildProjectCmd = &cobra.Command{
// 	Use:   "build",
// 	Args:  cobra.ExactArgs(0),
// 	Short: "build Docker image for current project",
// 	Long:  "build a Docker image for the Runpod project in the current folder",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//parse project toml
// 		//build Dockerfile
// 		//base image: from toml
// 		//run setup.sh for system deps
// 		//pip install requirements
// 		//cmd: start handler
// 		//docker build
// 		//print next steps
// 	},
// }

func init() {
	NewProjectCmd.Flags().StringVarP(&projectName, "name", "n", "", "project name")
	NewProjectCmd.Flags().StringVarP(&modelName, "model", "m", "", "model name")
	NewProjectCmd.Flags().StringVarP(&modelType, "type", "t", "", "model typype")
	NewProjectCmd.Flags().BoolVarP(&initCurrentDir, "init", "i", false, "use the current directory as the project directory")
}
