package gengast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

// AddFxOptionsFromModuleName adds fx options to the main module file, given the module name
func AddFxOptionsFromModuleName(path string, projModName, modName string) (string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("error parsing file, %s %w", path, err)
	}

	importPackage(node, projModName, modName)

	// Traverse the AST and find the fx.Options call
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "Options" {
					x.Args = append(x.Args, []ast.Expr{
						ast.NewIdent(modName + ".Module"),
					}...)
				}
			}
		}
		return true
	})

	// Add the source code in buffer
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return "", fmt.Errorf("error formatting ast nodes, err: %w", err)
	}

	formattedCode := buf.String()
	providerToInsert := fmt.Sprintf("%v.Module,", modName)
	formattedCode = strings.Replace(formattedCode, providerToInsert, "\n\t"+providerToInsert, 1)

	return formattedCode, nil
}

func importPackage(node *ast.File, projModName, modName string) {
	path := filepath.Join(projModName, "domain", modName)
	importSpec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%v"`, path),
		},
	}

	importDecl := &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: token.Pos(1), // for grouping
		Specs:  []ast.Spec{importSpec},
	}

	// Check if there are existing imports, and if so, add to them
	found := false
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok && genDecl.Tok == token.IMPORT {
			genDecl.Specs = append(genDecl.Specs, importSpec)
			found = true
			break
		}
	}

	// If no import declaration exists, add the new one to Decls
	if !found {
		node.Decls = append([]ast.Decl{importDecl}, node.Decls...)
	}
}
