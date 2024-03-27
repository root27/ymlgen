package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Workflow struct {
	Name string `yaml:"name"`
	On   On     `yaml:"on"`
	Jobs Jobs   `yaml:"jobs"`
}

type On struct {
	Push Push `yaml:"push"`
}

type Push struct {
	Branches []string `yaml:"branches"`
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

type Env struct {
	ServerKey string `yaml:"SERVER_KEY,omitempty"`
}

func main() {
	workflow := Workflow{}

	// Gather user inputs
	fmt.Print("Enter your repository name: ")
	var repoName string
	fmt.Scanln(&repoName)

	fmt.Print("Enter the trigger event (e.g., push, pull_request): ")
	var triggerEvent string
	fmt.Scanln(&triggerEvent)

	fmt.Print("Enter the branch name to trigger the workflow (e.g., main, master): ")
	var onBranch string
	fmt.Scanln(&onBranch)

	// Populate the Workflow struct
	workflow.Name = "Deployment"
	workflow.On.Push.Branches = []string{onBranch}
	workflow.Jobs.Deploy.RunsOn = "ubuntu-latest"
	workflow.Jobs.Deploy.Steps = []Step{
		{
			Uses: "actions/checkout@v2",
		},
		{
			Name: "Deploy to server",
			Run:  `echo "$SERVER_KEY" > secret && chmod 600 secret && ssh -o StrictHostKeyChecking=no -i secret root@185.247.139.226 -p 8357 'ls -la'`,
		},
	}

	// Set environment variables
	workflow.Jobs.Deploy.Steps[0].Env = Env{
		ServerKey: "${{ secrets.SERVER_KEY }}",
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
