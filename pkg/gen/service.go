package gen

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	gengast "github.com/mukezhz/geng/pkg/utils/ast"
)

type ServiceGenerator struct {
	cfg *models.Service
	fs  embed.FS
}

func NewServiceGenerator(
	cfg *models.Service,
	fs embed.FS,
) *ServiceGenerator {
	return &ServiceGenerator{
		cfg: cfg,
		fs:  fs,
	}
}

// GetChoices gets service generation choices
func (g *ServiceGenerator) GetChoices() ([]string, error) {
	servicePath := filepath.Join(".", "templates", "wesionary", "service")
	servicePath = utils.IgnoreWindowsPath(servicePath)

	temps, err := utils.ListEmbDir(g.fs, servicePath)
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

// SimilarChoice get similar choices to given choice data
func (g *ServiceGenerator) SimilarChoice(choices []string) ([]string, error) {
	allChoices, err := g.GetChoices()
	if err != nil {
		return nil, err
	}

	retChoices := make([]string, 0)
	for _, sourceChoice := range choices {
		for _, dstChoice := range allChoices {
			if strings.Contains(dstChoice, sourceChoice) {
				retChoices = append(retChoices, dstChoice)
			}
		}
	}

	return retChoices, nil
}

func (g *ServiceGenerator) Generate() error {
	logger := pkg.GetLogger()

	modPath := filepath.Join(g.cfg.Directory, "pkg", "services", "module.go")

	if len(g.cfg.ServiceType) == 0 {
		return nil
	}

	choices, err := g.GetChoices()
	if err != nil {
		return fmt.Errorf("cant get service choices, %w", err)
	}

	selectedChoices := make([]string, 0)
	for _, c := range g.cfg.ServiceType {
		if utils.StrInList(choices, c) {
			selectedChoices = append(selectedChoices, c)
		} else {
			logger.Warnf("skipping invalid choice: %s", c)
		}
	}

	logger.Infof("generating service, choices: %#v", selectedChoices)

	if err := g.updateProviders(modPath, selectedChoices); err != nil {
		return err
	}

	if err := g.addServiceFile(selectedChoices); err != nil {
		return err
	}

	logger.Infof("updated service in %s", g.cfg.Directory)

	return nil
}

// addInfraFile adds service file from template
func (g *ServiceGenerator) addServiceFile(selectedChoices []string) error {
	for _, choice := range selectedChoices {
		sourcePath := filepath.Join(".", "templates", "wesionary", "service", choice+".tmpl")
		sourcePath = utils.IgnoreWindowsPath(sourcePath)

		destPath := filepath.Join(g.cfg.Directory, "pkg", "services", choice+".go")

		err := utils.ExecTemplate(g.fs, sourcePath, destPath, g.cfg)
		if err != nil {
			return fmt.Errorf("error generating file from template. %w", err)
		}
	}

	return nil
}

// updateProviders generate function declarations for selected services and update providers
func (g *ServiceGenerator) updateProviders(modPath string, selectedChoices []string) error {
	var functions []string
	for _, choice := range selectedChoices {
		funcPath := filepath.Join(".", "templates", "wesionary", "service", choice+".tmpl")
		funcPath = utils.IgnoreWindowsPath(funcPath)

		funcDecl, err := gengast.GetFunctionDeclarations(g.fs, funcPath)
		if err != nil {
			return fmt.Errorf("error generating function declarations, path: %s, err: %w", funcPath, err)
		}

		// get only those declarations with new
		decl := make([]string, 0)
		for _, d := range funcDecl {
			if strings.Contains(d, "New") {
				decl = append(decl, d)
			}
		}

		functions = append(functions, decl...)
	}

	providerCode, err := gengast.AddListOfProvideInFxOptions(modPath, functions)
	if err != nil {
		return fmt.Errorf("error generating providers list. %w", err)
	}

	if err := utils.WriteToPath(modPath, providerCode); err != nil {
		return fmt.Errorf("couldn't write to: %s, err: %w", modPath, err)
	}

	return nil
}
