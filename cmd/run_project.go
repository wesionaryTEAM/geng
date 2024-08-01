package cmd

import (
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/spf13/cobra"
)

var runProjectCmd = &cobra.Command{
	Use:   "run:project",
	Short: "Run the project",
	Long: `
Run the project:
Alias to "go run main.go"
	`,
	Args: cobra.MaximumNArgs(1),
	Run:  runProject,
}

func runProject(_ *cobra.Command, args []string) {
	program := "go"
	commands := []string{"run", "main.go"}
	// execute command from golang
	err := utils.ExecuteCommand(program, commands, args...)
	if err != nil {
		return
	}
}
