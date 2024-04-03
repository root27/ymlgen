package main

import (
	"fmt"

	"github.com/fatih/color"

	parse "github.com/root27/yml-parser/editParser"
)

func Init() {

	green := color.New(color.FgGreen).SprintFunc()

	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println(green("Welcome to the GitHub Actions Workflow Generator"))

	welcome_message := `
		- This tool will help you create a GitHub Actions workflow file for your repository.
		
		- You will be prompted to enter the workflow name, trigger event, branch name, and the number of steps you would like to add to the workflow.
		
		- You will also be prompted to enter the uses, name, and run command for each step.
		
		- You can also add a secret to the workflow if you wish.
		
		- You can also edit the workflow file after it has been generated.
		`

	ps := color.New(color.FgRed).Add(color.Underline).PrintlnFunc()

	fmt.Println(cyan(welcome_message))

	ps("P.S.: If you use secret in your workflow, make sure to add the secret to your repository with the same name. Also you have to add run command to use the secret in your workflow.\n")

}

func main() {

	// init message

	Init()

	parse.Generator()

}
