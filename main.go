package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/root27/yml-parser/editParser"
	"github.com/root27/yml-parser/structs"
	"gopkg.in/yaml.v2"
)

func main() {

	// Gather user inputs
	fmt.Print("Enter the workflow name: ")
	var workFlowName string
	fmt.Scanln(&workFlowName)

	fmt.Print("Enter the trigger event (e.g., push, pull_request): ")
	var triggerEvent string
	fmt.Scanln(&triggerEvent)

	fmt.Print("Enter the branch name to trigger the workflow (e.g., main, master): ")
	var onBranch string
	fmt.Scanln(&onBranch)

	fmt.Printf("How many steps would you like to add to the workflow?: ")
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

		fmt.Printf("Enter the uses for step %d: ", i+1)
		var uses string
		fmt.Scanln(&uses)

		// Gather user inputs for steps
		fmt.Printf("Enter the step %d name: ", i+1)
		var stepName string
		fmt.Scanln(&stepName)

		fmt.Print("Would you like to add a secret? (y/n): ")

		fmt.Scanln(&addSecret)

		if strings.ToLower(addSecret) == "y" {

			fmt.Print("Enter the secret name: ")
			fmt.Scanln(&secretName)

			fmt.Printf("Enter the step %d command to run: ", i+1)

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
