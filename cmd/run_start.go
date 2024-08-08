package cmd

import (
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/spf13/cobra"
)

var runStartCmd = &cobra.Command{
	Use:   "run:start",
	Short: "Run the project execution cmd",
	Long: `
Execute the project:
Alias to "go run main.go app:serve"

For available command: see "geng project"
	`,
	Args: cobra.MaximumNArgs(1),
	Run:  startProject,
}

func startProject(_ *cobra.Command, args []string) {
	program := "go"
	if len(args) == 0 {
		args = append(args, "app:serve")
	}
	commands := []string{"run", "main.go"}
	// execute command from golang
	err := utils.ExecuteCommand(program, commands, args...)
	if err != nil {
		return
	}
}
