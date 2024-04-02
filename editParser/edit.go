package editParser

import (
	"fmt"
	"strings"

	"github.com/root27/yml-parser/structs"
)

func EditParser(yamlData []byte, workflow *structs.Workflow) error {

	lines := strings.Split(string(yamlData), "\n")

	for _, step := range workflow.Jobs.Deploy.Steps {

		if step.Env != nil && step.Run != "" {

			for i, line := range lines {

				if strings.Contains(line, "name: "+step.Name) && i != 0 {

					count := 0

					for _, char := range lines[i] {

						if char == ' ' {

							count++

						}

					}

					// fmt.Printf("Count: %d\n", count)

					lines[i] = strings.Repeat(" ", count-3) + "- name: " + step.Name

				}
			}

		}

	}

	newFileData := strings.Join(lines, "\n")

	fmt.Println(newFileData)

	return nil

}
