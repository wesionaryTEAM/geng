package gen

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
)

type InfraGenerator struct {
	cfg *models.Infrastructure
	fs  embed.FS
}

func NewInfraGenerator(
	cfg *models.Infrastructure,
	fs embed.FS,
) *InfraGenerator {
	return &InfraGenerator{
		cfg: cfg,
		fs:  fs,
	}
}

func (g *InfraGenerator) GetChoices() ([]string, error) {
	infraPath := filepath.Join(".", "templates", "wesionary", "infrastructure")
	infraPath = utils.IgnoreWindowsPath(infraPath)

	temps, err := utils.ListEmbDir(g.fs, infraPath)
	if err != nil {
		return nil, err
	}

	choices := make([]string, 0)
	for _, temp := range temps {
		choice := strings.Replace(temp, ".tmpl", "", 1)
		choices = append(choices, choice)
	}

	return choices, nil
}

func (g *InfraGenerator) Generate() error {
	logger := pkg.GetLogger()

	modPath := filepath.Join(g.cfg.Directory, "pkg", "infrastructure", "module.go")

	infraPath := filepath.Join(".", "templates", "wesionary", "infrastructure")
	infraPath = utils.IgnoreWindowsPath(infraPath)

	if len(g.cfg.InfraType) == 0 {
		return nil
	}

	choices, err := g.GetChoices()
	if err != nil {
		return fmt.Errorf("cant get infra choices, %w", err)
	}

	selectedChoices := make([]string, 0)
	for _, c := range g.cfg.InfraType {
		if utils.StrInList(choices, c) {
			selectedChoices = append(selectedChoices, c)
		} else {
			logger.Warnf("skipping invalid choice: %s", c)
		}
	}

	logger.Infof("generating infra, choices: %#v", selectedChoices)

	if err := g.updateProviders(modPath, selectedChoices); err != nil {
		return err
	}

	if err := g.addInfraFile(selectedChoices); err != nil {
		return err
	}

	logger.Infof("updated infrastructure in %s", g.cfg.Directory)

  if err := g.addSimilarServices(selectedChoices); err != nil {
    return err
  }

	return nil
}

func (g *InfraGenerator) addSimilarServices(selectedChoices []string) error {
	serviceGen := ServiceGenerator{
		fs: g.fs,
		cfg: &models.Service{
			Directory:         g.cfg.Directory,
			ProjectModuleName: g.cfg.ProjectModuleName,
		},
	}
	services, err := serviceGen.SimilarChoice(selectedChoices)
	if err != nil {
		return fmt.Errorf("cant get similar choices as infrastructure")
	}

	serviceGen.cfg.ServiceType = services
	if err := serviceGen.Generate(); err != nil {
		return err
	}

  return nil
}

// addInfraFile adds infrastructure file from template
func (g *InfraGenerator) addInfraFile(selectedChoices []string) error {
	for _, choice := range selectedChoices {
		sourcePath := filepath.Join(".", "templates", "wesionary", "infrastructure", choice+".tmpl")
		sourcePath = utils.IgnoreWindowsPath(sourcePath)

		destPath := filepath.Join(g.cfg.Directory, "pkg", "infrastructure", choice+".go")

		err := utils.ExecTemplate(g.fs, sourcePath, destPath, g.cfg)
		if err != nil {
			return fmt.Errorf("error generating file from template. %w", err)
		}
	}

	return nil
}

// updateProviders generate function declarations for selected infrastructure and update providers
func (g *InfraGenerator) updateProviders(modPath string, selectedChoices []string) error {
	var functions []string
	for _, choice := range selectedChoices {
		funcPath := filepath.Join(".", "templates", "wesionary", "infrastructure", choice+".tmpl")
		funcPath = utils.IgnoreWindowsPath(funcPath)

		funcDecl, err := utils.GetFunctionDeclarations(g.fs, funcPath)
		if err != nil {
			return fmt.Errorf("error generating function declarations, path: %s, err: %w", funcPath, err)
		}
		functions = append(functions, funcDecl...)
	}

	providerCode, err := utils.AddListOfProvideInFxOptions(modPath, functions)
	if err != nil {
		return fmt.Errorf("error generating providers list. %w", err)
	}

	if err := utils.WriteToPath(modPath, providerCode); err != nil {
		return fmt.Errorf("couldn't write to: %s, err: %w", modPath, err)
	}

	return nil
}
