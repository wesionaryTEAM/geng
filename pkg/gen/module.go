package gen

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	gengast "github.com/mukezhz/geng/pkg/utils/ast"
)

type ModuleGenerator struct {
	cfg *models.Module
	fs  embed.FS
}

func NewModuleGenerator(cfg *models.Module, fs embed.FS) *ModuleGenerator {
	return &ModuleGenerator{cfg: cfg, fs: fs}
}

func (g *ModuleGenerator) Generate() error {

	mainModPath := filepath.Join(g.cfg.Directory, "domain", "module.go")
	targetRoot := filepath.Join(g.cfg.Directory, "domain", g.cfg.PackageName)

	templatePath := filepath.Join(".", "templates", "wesionary", "module")
	templatePath = utils.IgnoreWindowsPath(templatePath)

	files, err := utils.GenerateFiles(g.fs, templatePath, targetRoot, g.cfg)
	if err != nil {
		return fmt.Errorf("error generating module files: %w", err)
	}

	op, err := gengast.AddFxOptionsFromModuleName(mainModPath, g.cfg.ProjectModuleName, g.cfg.PackageName)
	if err != nil {
		return fmt.Errorf("error adding fx options to main module: %w", err)
	}

	if err := utils.WriteToPath(mainModPath, op); err != nil {
		return fmt.Errorf("error writing to main module: %w", err)
	}

	logger := pkg.GetLogger()
	for _, file := range files {
		logger.Infof("generated: %s", file)
	}

	logger.Infof("updated: %s", mainModPath)

	return nil
}
