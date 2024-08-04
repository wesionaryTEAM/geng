package gengast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

// AddListOfProvideInFxOptions adds provide in fx.Options list in path
func AddListOfProvideInFxOptions(path string, providerList []string) (string, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}

	// Track existing providers
	existingProviders := make(map[string]any)

	// Traverse the AST and find the fx.Options call
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Options" {

				// Check existing arguments in fx.Options
				for _, arg := range x.Args {
					if callExpr, ok := arg.(*ast.CallExpr); ok {
						if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok && selExpr.Sel.Name == "Provide" {
							if len(callExpr.Args) > 0 {
								if ident, ok := callExpr.Args[0].(*ast.Ident); ok {
									existingProviders[ident.Name] = struct{}{}
								}
							}
						}
					}
				}

				// Add new providers
				for _, provider := range providerList {
					if _, exists := existingProviders[provider]; !exists {
						x.Args = append(x.Args, &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fx"),
								Sel: ast.NewIdent("Provide"),
							},
							Args: []ast.Expr{
								ast.NewIdent(provider),
							},
							Rparen: token.Pos(1),
						})
					}
				}
			}
		}
		return true
	})

	// Add the source code in buffer
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return "", fmt.Errorf("node formatting error. %w", err)
	}

	formattedCode := buf.String()
	for _, provider := range providerList {
		if _, exists := existingProviders[provider]; exists {
			continue
		}

		providerToInsert := fmt.Sprintf("fx.Provide(%v)", provider)
		formattedCode = strings.Replace(formattedCode, providerToInsert, "\n\t\t"+providerToInsert, -1)
	}

	return formattedCode, nil
}
