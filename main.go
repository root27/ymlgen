package main

import (
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
	RunsOn string      `yaml:"runs-on"`
	Steps  interface{} `yaml:"steps"`
}

type Step struct {
	Uses string      `yaml:"uses,omitempty"`
	Name string      `yaml:"name,omitempty"`
	Env  interface{} `yaml:"env"`
	Run  string      `yaml:"run,omitempty"`
}

// type Env struct {
// 	ServerKey string `yaml:"SERVER_KEY,omitempty"`
// }

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

	if strings.ToLower(addSecret) == "y" {
		fmt.Print("Enter the secret name: ")
		var secretName string
		fmt.Scanln(&secretName)

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
	// workflow.Jobs.Deploy.Steps = []Step{
	// 	{
	// 		Uses: "actions/checkout@v2",
	// 	},
	// 	{
	// 		Name: "Deploy to server",
	// 		Env: map[string]interface{}{
	// 			secretName: "${{ secrets." + secretName + " }}",
	// 		},
	// 		Run: `echo "$SERVER_KEY" > secret && chmod 600 secret && ssh -o StrictHostKeyChecking=no -i secret root@185.247.139.226 -p 8357 'ls -la'`,
	// 	},
	// }

	workflow.Jobs.Deploy.Steps = map[string]interface{}{
		"uses": "actions/checkout@v2",
		"name": "Deploy to server",
		"env": map[string]interface{}{
			secretName: "${{ secrets." + secretName + " }}",
		},
	}

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
