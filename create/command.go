package create

import (
	"os"
	"os/exec"
)

func RunCommand(c Command) {
	command := exec.Command(c.Name, c.Args...)
	if c.Dir != nil {
		dir := *c.Dir
		os.MkdirAll(dir, os.ModePerm)
		command.Dir = dir
	}

	// Run the command and pass the console to it
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
	command.Wait()
}
