package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geng",
	Short: "A golang project and module generator cli tool",
	Long:  "geng is a CLI tool for golang api project generation.",
}

var rootFlags = []utils.FlagItem{
	{Name: "dir", Short: "d", Desc: "target directory for geng", Persistent: true},
	{Name: "goversion", Short: "v", Desc: "golang version in mod file", Persistent: true},
}

func init() {
	utils.SetFlags(rootCmd, rootFlags)
	utils.BindFlag(rootCmd, rootFlags)

	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(infraCmd)
	rootCmd.AddCommand(serviceCmd)
}

func Execute() {
	logger := pkg.GetLogger()

	pkg.PrintIntro()

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("couldn't execute the necessary command", "err", err)
	}
}
