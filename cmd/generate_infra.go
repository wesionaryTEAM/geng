package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
	"github.com/spf13/cobra"
)

var infraCmd = &cobra.Command{
	Use:   "add:infra",
	Short: "Adds infrastructure modules in project",
	Run:   createInfra,
}

var inraFlags = []utils.FlagItem{
	{Name: "mod", Short: "m", Desc: "module name for the project"},
	{Name: "infratype", Short: "i", Desc: "infastructure types (eg; aws, firebase, gcp, gmail, opensearch)", Def: []string{}},
}

func init() {
	utils.SetFlags(infraCmd, inraFlags)
}

func createInfra(cmd *cobra.Command, args []string) {
	utils.BindFlag(cmd, inraFlags)

	logger := pkg.GetLogger()

	input := pkg.GetConfig[models.Infrastructure]()

	// Fill up data items, if can be filled up automatically
	input.AutoFill()

	// TODO: show terminal ui, auto fill up contents already given by flags
	logger.Infof("input: %#v", input)

	// validate input data
	if err := input.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	generator := gen.NewInfraGenerator(input, templates.FS)
	if err := generator.Generate(); err != nil {
		logger.Fatal("project generation failed.", "err", err)
	}

}
