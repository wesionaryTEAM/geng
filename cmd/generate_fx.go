package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/spf13/cobra"
)

var fxCmd = &cobra.Command{
	Use:   "fx",
	Short: "Generate the fx configuration file for the project",
	Long: `
Generate a module by reading comments from the source code.
Example: 
  geng fx

	`,
	Run: generateFx,
}

func generateFx(_ *cobra.Command, _ []string) {

	logger := pkg.GetLogger()

	cfg := pkg.GetConfig[models.Fx]()
	cfg.AutoFill()

	if err := cfg.Validate(); err != nil {
		logger.Fatal("config validation failed", "err", err)
	}

	fxGenerator := gen.NewFxGenerator(cfg)
	if err := fxGenerator.Generate(); err != nil {
		logger.Fatal("couldn't generate providers from comments", "err", err)
	}
}
