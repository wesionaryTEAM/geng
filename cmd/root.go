package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geng",
	Short: "A golang project and module generator cli tool",
	Long:  "geng is a CLI tool for golang api project generation.",
}

func init() {
	pkg.BindStrPFlag(rootCmd, "dir", "d", "", "target directory to create project at")
	pkg.BindStrPFlag(rootCmd, "goversion", "v", "", "go version for project generation")

	rootCmd.AddCommand(projectCmd)
}

func Execute() {
	logger := pkg.GetLogger()

	pkg.PrintIntro()

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("couldn't execute the necessary command", "err", err)
	}
}
