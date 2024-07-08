package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/timmo001/create-timmo-app/create"
)

type promptItem struct {
	Label         string
	Value         string
	PromptSubType *promptui.Select
}

var itemsTypeTypescriptBasic = []promptItem{
	{Label: "Bun", Value: "bun"},
	{Label: "Node.js w/ yarn", Value: "node-yarn"},
}

var itemsTypeTypescript = []promptItem{
	{Label: "Basic", Value: "basic", PromptSubType: &promptui.Select{
		Label: "What type of basic TypeScript application do you want to create?",
		Items: itemsTypeTypescriptBasic,
	}},
	{Label: "T3", Value: "t3"},
	{Label: "Next.js", Value: "nextjs"},
}

var itemsTypeGo = []promptItem{
	{Label: "CLI", Value: "cli"},
}

var itemsType = []promptItem{
	{Label: "TypeScript", Value: "typescript", PromptSubType: &promptui.Select{
		Label: "What type of TypeScript application do you want to create?",
		Items: itemsTypeTypescript,
	}},
	{Label: "Go", Value: "go", PromptSubType: &promptui.Select{
		Label: "What type of Go application do you want to create?",
		Items: itemsTypeGo,
	}},
}

var promptAppType = promptui.Select{
	Label: "What type of application do you want to create?",
	Items: itemsType,
}

var promptName = promptui.Prompt{
	Label:    "What is the name of your application?",
	Validate: validateName,
}

func Prompt() {
	// Run type prompts
	appType, err := processSelect(&promptAppType)
	if err != nil {
		return
	}

	// Run name prompt
	appName, err := promptName.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the results
	fmt.Println("Options:")
	fmt.Printf("  Type: %s\n", appType)
	fmt.Printf("  Name: %s\n", appName)

	// Create the application
	create.CreateApp(create.App{
		Type: appType,
		Name: appName,
	})
}

func processSelect(promptSelect *promptui.Select) (string, error) {
	i, _, err := promptSelect.Run()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Check if the item has a sub type prompt
	items := promptSelect.Items.([]promptItem)

	if items[i].PromptSubType != nil {
		// Run the sub type prompt
		subType, err := processSelect(items[i].PromptSubType)
		if err != nil {
			return "", err
		}

		fmt.Printf("%s-%s", items[i].Value, subType)

		return fmt.Sprintf("%s-%s", items[i].Value, subType), nil
	}

	return items[i].Value, nil
}

func validateName(input string) error {
	// Check the name is not empty
	if input == "" {
		return fmt.Errorf("the name of the application cannot be empty")
	}

	// Check the name is a valid directory name

	// Check the name is not already in use

	return nil
}
