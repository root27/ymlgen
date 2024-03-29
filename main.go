package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Workflow struct {
	Name string      `yaml:"name"`
	On   interface{} `yaml:"on"`
	Jobs Jobs        `yaml:"jobs"`
}

type Jobs struct {
	Deploy Deploy `yaml:"deploy"`
}

type Deploy struct {
	RunsOn string `yaml:"runs-on"`
	Steps  []Step `yaml:"steps"`
}

type Step struct {
	Uses string `yaml:"uses,omitempty"`
	Name string `yaml:"name,omitempty"`
	Env  Env    `yaml:"env,omitempty"`
	Run  string `yaml:"run,omitempty"`
}

type Env map[string]interface{}

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

	fmt.Print("Would you like to add a secret? (y/n): ")
	var addSecret string
	fmt.Scanln(&addSecret)

	var secretName string

	var runCommand string

	if strings.ToLower(addSecret) == "y" {
		fmt.Print("Enter the secret name: ")

		fmt.Scanln(&secretName)

		reader := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter the command to run: ")

		reader.Scan()

		runCommand = reader.Text()

	}

	// Populate the Workflow struct

	workflow := Workflow{}

	workflow.Name = workFlowName

	workflow.On = map[string]interface{}{
		triggerEvent: map[string]interface{}{
			"branches": []string{onBranch},
		},
	}

	workflow.Jobs.Deploy.RunsOn = "ubuntu-latest"

	workflow.Jobs.Deploy.Steps = []Step{
		{
			Uses: "actions/checkout@v2",
		},
		{
			Name: "Test Action",
		},
	}

	if secretName != "" {
		workflow.Jobs.Deploy.Steps[1].Env = Env{
			secretName: "${{ secrets." + secretName + " }}",
		}

		workflow.Jobs.Deploy.Steps[1].Run = runCommand

	}

	fmt.Printf("Workflow steps: %v\n", workflow.Jobs.Deploy.Steps)

	// Convert struct to YAML
	yamlData, err := yaml.Marshal(&workflow)
	if err != nil {
		fmt.Printf("Error marshalling YAML: %v\n", err)
		return
	}

	// Write YAML to file
	err = os.WriteFile(".github/workflows/deployment.yml", yamlData, 0644)
	if err != nil {
		fmt.Printf("Error writing YAML file: %v\n", err)
		return
	}

	fmt.Println("Deployment YAML file generated successfully.")
}
