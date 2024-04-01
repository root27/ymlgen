package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/root27/yml-parser/editParser"
	"github.com/root27/yml-parser/structs"
	"gopkg.in/yaml.v2"

	"github.com/fatih/color"
)

func Init() {

	green := color.New(color.FgGreen).SprintFunc()

	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println(green("Welcome to the GitHub Actions Workflow Generator"))

	welcome_message := `
		This tool will help you create a GitHub Actions workflow file for your repository.
		You will be prompted to enter the workflow name, trigger event, branch name, and the number of steps you would like to add to the workflow.
		You will also be prompted to enter the uses, name, and run command for each step.
		You can also add a secret to the workflow if you wish.
		You can also edit the workflow file after it has been generated.
		`

	fmt.Println(cyan(welcome_message))

}

func main() {

	// init message

	Init()

	titleGreen := color.New(color.FgGreen).Add(color.Underline).SprintFunc()

	// Gather user inputs
	fmt.Print(titleGreen("Enter the workflow name: "))
	var workFlowName string
	fmt.Scanln(&workFlowName)

	fmt.Print(titleGreen("Enter the trigger event (e.g., push, pull_request): "))
	var triggerEvent string
	fmt.Scanln(&triggerEvent)

	fmt.Print(titleGreen("Enter the branch name to trigger the workflow (e.g., main, master): "))
	var onBranch string
	fmt.Scanln(&onBranch)

	fmt.Print(titleGreen("How many steps would you like to add to the workflow?: "))
	var stepCount int
	fmt.Scanln(&stepCount)

	var addSecret string

	var secretName string

	var runCommand string

	// Populate the Workflow struct

	workflow := structs.Workflow{}

	workflow.Name = workFlowName

	workflow.On = map[string]interface{}{
		triggerEvent: map[string]interface{}{
			"branches": []string{onBranch},
		},
	}

	workflow.Jobs.Deploy.RunsOn = "ubuntu-latest"

	for i := 0; i < stepCount; i++ {

		usesString := color.New(color.FgGreen).Sprintf("Enter the uses for step %d (e.g., actions/checkout): ", i+1)
		fmt.Print(usesString)
		var uses string
		fmt.Scanln(&uses)

		// Gather user inputs for steps

		stepNameString := color.New(color.FgGreen).Sprintf("Enter the step %d name: ", i+1)
		fmt.Print(stepNameString)
		var stepName string
		fmt.Scanln(&stepName)

		fmt.Print(titleGreen("Would you like to add a secret? (y/n): "))

		fmt.Scanln(&addSecret)

		if strings.ToLower(addSecret) == "y" {

			fmt.Print(titleGreen("Enter the secret name: "))
			fmt.Scanln(&secretName)

			commandString := color.New(color.FgGreen).Sprintf("Enter the run command for step %d (e.g., echo 'Hello World'): ", i+1)
			fmt.Print(commandString)
			scanner := bufio.NewScanner(os.Stdin)

			scanner.Scan()

			runCommand = scanner.Text()

			// Add step to workflow

			workflow.Jobs.Deploy.Steps = append(workflow.Jobs.Deploy.Steps, structs.Step{
				Uses: uses,
				Name: stepName,
				Env: structs.Env{
					secretName: "${{ secrets." + secretName + " }}",
				},
				Run: runCommand,
			})

			continue
		}

		// Add step to workflow

		workflow.Jobs.Deploy.Steps = append(workflow.Jobs.Deploy.Steps, structs.Step{
			Uses: uses,
			Name: stepName,
		})

	}

	// Convert struct to YAML
	yamlData, err := yaml.Marshal(&workflow)
	if err != nil {
		fmt.Printf("Error marshalling YAML: %v\n", err)
		return
	}

	err = editParser.EditParser(yamlData, &workflow)

	if err != nil {
		fmt.Printf("Error editing YAML: %v\n", err)
		return
	}

	fmt.Println("Workflow created successfully")

}
