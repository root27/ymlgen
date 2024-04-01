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

	parse.Generator()

}
