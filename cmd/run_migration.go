package cmd

import (
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/spf13/cobra"
)

var runMigrateCmd = &cobra.Command{
	Use:   "run:migrate",
	Short: "run the migrate command",
	Args:  cobra.MaximumNArgs(0),
	Run:   migrationProject,
}

func migrationProject(_ *cobra.Command, args []string) {
	program := "go"
	commands := []string{"run", "main.go", "migrate:run"}

	// execute command from golang
	err := utils.ExecuteCommand(program, commands, args...)
	if err != nil {
		return
	}
}
