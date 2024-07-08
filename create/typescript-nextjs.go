package create

func CreateNextApp(app App) {
	RunCommand(Command{
		Name: "bun",
		Args: []string{"create", "next-app@latest",
			"--typescript", "--tailwind", "--eslint", "--app", "--src-dir",
			"--import-alias", "~/*", "--use-bun", app.Name},
	})

	// TODO: Replace files (README.md, .github/*, etc.)
}
