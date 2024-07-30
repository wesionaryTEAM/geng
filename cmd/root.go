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

func ExecuteWith(cfg *pkg.GengConfig) {
	logger := pkg.GetLogger()

  pkg.PrintIntro()
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("couldn't execute the necessary command", "err", err)
	}
}
