package main

import (
	"fmt"
	"os/exec"
)

type App struct {
	Type string
	Name string
}

func CreateApp(app App) {
	fmt.Printf("Creating a new application of type %s with the name %s\n", app.Type, app.Name)

	// Create the application
	switch app.Type {
	case "typescript-basic-bun":
		exec.Command("mkdir", app.Name).Start()
		exec.Command("cd", app.Name).Start()
		exec.Command("bun", "init").Start()
	case "typescript-basic-node-yarn":
		exec.Command("mkdir", app.Name).Start()
		exec.Command("cd", app.Name).Start()
		exec.Command("yarn", "init").Start()
	case "typescript-t3":
		exec.Command("bun", "create-t3-app@latest", app.Name).Start()
	default:
		fmt.Printf("Unknown type selected: %s", app.Type)
	}

}
