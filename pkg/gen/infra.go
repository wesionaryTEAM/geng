package gen

import (
	"embed"
	"path/filepath"

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

func (g *InfraGenerator) Generate() error {
	logger := pkg.GetLogger()

	modPath := filepath.Join(g.cfg.Directory, "pkg", "infrastructure", "module.go")

	infraPath := filepath.Join(".", "templates", "wesionary", "infrastructure")
	infraPath = utils.IgnoreWindowsPath(infraPath)

	logger.Info(modPath, infraPath)
	return nil
}
