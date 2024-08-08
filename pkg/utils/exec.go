package utils

import (
	"os"
	"os/exec"
)

// ExecuteCommand executes a command
func ExecuteCommand(program string, commands []string, args ...string) error {
	var cmd *exec.Cmd
	runCommand := commands
	if len(args) == 0 {
		cmd = exec.Command(program, runCommand...)
	} else {
		runCommand = append(runCommand, args...)
		cmd = exec.Command(program, runCommand...)
	}

	cmd.Dir = "." // Ensure this is correct

	// Set Stdout and Stderr to os.Stdout, os.Stderr respectively
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command and check for errors
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
