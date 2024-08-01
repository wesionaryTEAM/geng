package cmd

import (
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/spf13/cobra"
)

var runSeedCmd = &cobra.Command{
	Use:   "run:seed",
	Short: "Run the Seed command",
	Run:   seedProject,
}

func seedProject(_ *cobra.Command, args []string) {
	program := "go"
	commands := []string{"run", "main.go", "seed:run"}

	// execute command from golang
	if len(args) == 0 || (len(args) == 1 && args[0] == "all") {
		commands = append(commands, "--all")
	} else {
		for _, arg := range args {
			commands = append(commands, "--name")
			commands = append(commands, arg)
		}
	}
	err := utils.ExecuteCommand(program, commands, args...)
	if err != nil {
		return
	}
}
