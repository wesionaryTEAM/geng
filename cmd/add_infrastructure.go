package cmd

import (
	"embed"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/mukezhz/geng/pkg/constant"
	"github.com/mukezhz/geng/pkg/model"
	"github.com/mukezhz/geng/pkg/terminal"
	"github.com/mukezhz/geng/pkg/utility"
	"github.com/spf13/cobra"
)

var addInfrastructureCmd = &cobra.Command{
	Use:   "infra add [name]",
	Short: "Add a new infrastructure",
	Args:  cobra.MaximumNArgs(2),
	Run:   addInfrastructureHandler,
}

func addInfrastructureHandler(_ *cobra.Command, args []string) {
	if len(args) > 0 && !strings.Contains(args[0], "add") {
		color.Redln("Error: invalid command")
		return
	}
	projectModule, err := utility.GetModuleNameFromGoModFile()
	if err != nil {
		color.Redln("Error finding Module name from go.mod:", err)
		return
	}
	data := utility.GetModuleDataFromModuleName("", projectModule.Module, projectModule.GoVersion)
	currentDir, err := os.Getwd()
	if err != nil {
		color.Redln("Error getting current directory:", err)
		panic(err)
	}
	projectPath, err := utility.FindGitRoot(currentDir)
	if err != nil {
		color.Redln("Error finding Git root:", err)
		return
	}
	infrastructureModulePath := filepath.Join(projectPath, "pkg", "infrastructure", "module.go")
	templateInfraPath := filepath.Join(".", "templates", "wesionary", "infrastructure")
	infrasTmpl := utility.ListDirectory(templatesFS, templateInfraPath)
	infras := utility.Map[string, string](infrasTmpl, func(q string) string {
		return strings.Replace(q, ".tmpl", "", 1)
	})
	questions := []terminal.ProjectQuestion{
		terminal.NewCheckboxQuestion(constant.InfrastructureNameKEY, "Select the infrastructure? [<space> to select]", infras),
	}

	terminal.StartInteractiveTerminal(questions)

	items := addInfrastructure(questions, infrasTmpl, infrastructureModulePath, data, false, templatesFS)
	if len(items) == 0 {
		color.Red.Println("No infrastructure selected")
		return
	}
	utility.PrintColorizeInfrastructureDetail(data, infras)
}

func addInfrastructure(
	questions []terminal.ProjectQuestion,
	infrasTmpl []string,
	infrastructureModulePath string,
	data model.ModuleData,
	isNewProject bool,
	templatesFS embed.FS) []int {
	var functions []string
	var items []int
	for _, q := range questions {
		switch q.Key {
		case constant.InfrastructureNameKEY:
			selected := q.Input.Selected()
			for s := range selected {
				functions = append(functions,
					utility.GetFunctionDeclarations(filepath.Join(".", "templates", "wesionary", "infrastructure", infrasTmpl[s]), templatesFS)...,
				)
				items = append(items, s)
			}
		}
	}
	if len(items) == 0 {
		return items
	}
	updatedCode := utility.AddListOfProvideInFxOptions(infrastructureModulePath, functions)
	utility.WriteContentToPath(infrastructureModulePath, updatedCode)

	var servicesTmpl []string
	for _, i := range items {
		templatePath := filepath.Join(".", "templates", "wesionary", "infrastructure", infrasTmpl[i])
		var targetRoot string
		if isNewProject {
			targetRoot = filepath.Join(data.PackageName, "pkg", "infrastructure", strings.Replace(infrasTmpl[i], ".tmpl", ".go", 1))
		} else {
			targetRoot = filepath.Join(".", "pkg", "infrastructure", strings.Replace(infrasTmpl[i], ".tmpl", ".go", 1))
		}

		fileName := strings.Replace(infrasTmpl[i], ".tmpl", "", 1)
		serviceTemplatePath := filepath.Join(".", "templates", "wesionary", "service")
		for _, file := range utility.ListDirectory(templatesFS, serviceTemplatePath) {
			if strings.Contains(file, fileName) {
				servicesTmpl = append(servicesTmpl, file)
			}
		}
		utility.GenerateFromEmbeddedTemplate(templatesFS, templatePath, targetRoot, data)

		// add service too
		var serviceModulePath string
		if isNewProject {
			serviceModulePath = filepath.Join(data.PackageName, "pkg", "services", "module.go")
		} else {
			serviceModulePath = filepath.Join(".", "pkg", "services", "module.go")
		}

		addService(questions, servicesTmpl, serviceModulePath, data, true, templatesFS)
	}
	return items
}
