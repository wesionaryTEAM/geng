package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "add:service",
	Short: "Adds service modules in project",
	Run:   createService,
}

var serviceFlags = []utils.FlagItem{
	{Name: "mod", Short: "m", Desc: "module name for the project"},
	{Name: "servicetype", Short: "s", Desc: "service types (eg; aws, firebase, gcp, gmail, opensearch)", Def: []string{}},
}

func init() {
	utils.SetFlags(serviceCmd, serviceFlags)
}

func createService(cmd *cobra.Command, args []string) {
	utils.BindFlag(cmd, serviceFlags)

	logger := pkg.GetLogger()

	input := pkg.GetConfig[models.Service]()

	// Fill up data items, if can be filled up automatically
	input.AutoFill()

	// TODO: show terminal ui, auto fill up contents already given by flags
	logger.Infof("input: %#v", input)

	// validate input data
	if err := input.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	generator := gen.NewServiceGenerator(input, templates.FS)
	if err := generator.Generate(); err != nil {
		logger.Fatal("project generation failed.", "err", err)
	}

}
