package gengast

import (
	"embed"
	"go/ast"
	"go/parser"
	"go/token"
)

// GetFunctionDeclarations gets all the function declarations present in a go file present in embed
func GetFunctionDeclarations(fs embed.FS, path string) ([]string, error) {
	fset := token.NewFileSet()
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}

	node, err := parser.ParseFile(fset, "", f, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var functions []string
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			functions = append(functions, x.Name.Name)
		}
		return true
	})

	return functions, nil
}
