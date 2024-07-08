package create

import (
	"fmt"
)

type App struct {
	Type string
	Name string
}

type Command struct {
	Name string
	Args []string
	Dir  *string
}

func CreateApp(app App) {
	fmt.Printf("Creating a new application of type %s with the name %s\n", app.Type, app.Name)

	// Create the application
	switch app.Type {
	case "typescript-basic-bun":
		RunCommand(Command{
			Name: "bun",
			Args: []string{"init", "-y"},
			Dir:  &app.Name,
		})
	case "typescript-basic-node-yarn":
		RunCommand(Command{
			Name: "yarn",
			Args: []string{"init", "-y"},
			Dir:  &app.Name,
		})
	case "typescript-t3":
		RunCommand(Command{
			Name: "bun",
			Args: []string{"create", "t3-app@latest", app.Name},
		})
	case "typescript-nextjs":
		CreateNextApp(app)
	default:
		fmt.Printf("Unknown type selected: %s", app.Type)
	}

}

