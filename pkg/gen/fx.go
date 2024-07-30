package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FxGenerator struct {
	cfg *models.Fx
}

func NewFxGenerator(cfg *models.Fx) *FxGenerator {
	return &FxGenerator{
		cfg: cfg,
	}
}

func (g *FxGenerator) Generate() error {
	fset := token.NewFileSet()
	root := g.cfg.Directory

	moduleMap := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse file %s, %w", path, err)
		}

		processedFile := g.processFile(file)
		if len(processedFile) != 0 {
			moduleMap[filepath.Dir(path)] = append(moduleMap[filepath.Dir(path)], processedFile...)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking through directory. %w", err)
	}

	count, err := g.generateFile(moduleMap)
	if err != nil {
		return fmt.Errorf("error generating modules. %w", err)
	}

	pkg.GetLogger().Infof("generated %d module files", count)

	return nil
}

func (g *FxGenerator) processFile(file *ast.File) []string {
	providers := make([]string, 0)

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Doc == nil {
			return true
		}

		var provide bool
		var asInterface string

		for _, comment := range fn.Doc.List {
			if strings.Contains(comment.Text, "@fxProvide") {
				provide = true
			}
			if strings.HasPrefix(comment.Text, "// @fxAs ") {
				parts := strings.SplitN(comment.Text, " ", 3)
				if len(parts) > 2 {
					asInterface = strings.TrimSpace(parts[2])
				}
			}
		}

		if provide {
			if asInterface != "" {
				providers = append(providers, fmt.Sprintf("fx.Provide(fx.Annotate(%s, fx.As(new(%s)))),", fn.Name.Name, asInterface))
			} else {
				providers = append(providers, fmt.Sprintf("fx.Provide(%s),", fn.Name.Name))
			}
		}
		return true

	})

	return providers
}

const templateText = `package {{.PackageName}}

import "go.uber.org/fx"

// This file is generated by geng. DO NOT EDIT.
// File generated at: {{.Timestamp}}
// File will be overwritten when running geng again.

var {{.ModuleName}} = fx.Module("{{.Name}}",
{{- range .Providers}}
    {{.}}
{{- end}}
)
`

func (g *FxGenerator) generateFile(deps map[string][]string) (int, error) {
	tmpl, err := template.New("module").Parse(templateText)
	if err != nil {
		return 0, fmt.Errorf("error parsing template. %w", err)
	}

	for k, v := range deps {
		module := map[string]interface{}{
			"PackageName": filepath.Base(k),
			"Name":        filepath.Base(k),
			"ModuleName":  cases.Title(language.English).String(filepath.Base(k)),
			"Providers":   v,
		}

		filePath := filepath.Join(k, "module.go")
		file, err := os.Create(filePath)
		if err != nil {
			return 0, fmt.Errorf("cannot create file. path: %s ,err: %w", filePath, err)
		}

		defer file.Close()

		err = tmpl.Execute(file, module)
		if err != nil {
			return 0, fmt.Errorf("error executing template. err: %w", err)
		}
	}

	return len(deps), nil
}