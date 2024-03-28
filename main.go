package main

import (
	"fmt"
	"reflect"
)

type Workflow struct {
	Name string `yaml:"name"`
	On   On     `yaml:"on"`
	Jobs Jobs   `yaml:"jobs"`
}

type On struct {
	Push Push `yaml:"push"`
	Pull Pull `yaml:"pull_request"`
}

type Pull struct {
	Branches []string `yaml:"branches"`
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
	fmt.Print("Enter the workflow name: ")
	var workFlowName string
	fmt.Scanln(&workFlowName)

	fmt.Print("Enter the trigger event (e.g., push, pull_request): ")
	var triggerEvent string
	fmt.Scanln(&triggerEvent)

	fmt.Print("Enter the branch name to trigger the workflow (e.g., main, master): ")
	var onBranch string
	fmt.Scanln(&onBranch)

	// Populate the Workflow struct

	flowType := reflect.TypeOf(workflow)

	newStruct := reflect.New(flowType)

	newStruct.Elem().FieldByName("Name").SetString(workFlowName)

	newStruct.Elem().FieldByName("On").FieldByName(triggerEvent).FieldByName("Branches").Set(reflect.ValueOf([]string{onBranch}))

	fmt.Print(newStruct.Interface())

	// workflow.On.Push.Branches = []string{onBranch}

	// workflow.Jobs.Deploy.RunsOn = "ubuntu-latest"
	// workflow.Jobs.Deploy.Steps = []Step{
	// 	{
	// 		Uses: "actions/checkout@v2",
	// 	},
	// 	{
	// 		Name: "Deploy to server",
	// 		Env: Env{
	// 			ServerKey: "${{ secrets.SERVER_KEY }}",
	// 		},
	// 		Run: `echo "$SERVER_KEY" > secret && chmod 600 secret && ssh -o StrictHostKeyChecking=no -i secret root@185.247.139.226 -p 8357 'ls -la'`,
	// 	},
	// }

	// // Convert struct to YAML
	// yamlData, err := yaml.Marshal(&workflow)
	// if err != nil {
	// 	fmt.Printf("Error marshalling YAML: %v\n", err)
	// 	return
	// }

	// // Write YAML to file
	// err = os.WriteFile(".github/workflows/deployment.yml", yamlData, 0644)
	// if err != nil {
	// 	fmt.Printf("Error writing YAML file: %v\n", err)
	// 	return
	// }

	// fmt.Println("Deployment YAML file generated successfully.")
}
