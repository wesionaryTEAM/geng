package gen

import (
	"embed"
	"path/filepath"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
)

type ProjectGenerator struct {
	cfg *models.Project
	fs  embed.FS
}

func NewProjectGenerator(cfg *models.Project, fs embed.FS) *ProjectGenerator {
	return &ProjectGenerator{
		cfg: cfg,
		fs:  templates.FS,
	}
}

func (g *ProjectGenerator) Generate() error {

	tempPath := filepath.Join("templates", "wesionary", "project")
	tempPath = utils.IgnoreWindowsPath(tempPath)

	generatedFiles, err := utils.GenerateFiles(g.fs, tempPath, g.cfg.Directory, g.cfg)
	if err != nil {
		return err
	}

	logger := pkg.GetLogger()
	logger.Info("project generation success")
	logger.Infof("generated %d files", len(generatedFiles))

	if err := utils.ModTidy(g.cfg.Directory); err != nil {
		return err
	}

	return nil
}
