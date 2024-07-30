package utils

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// GenerateFiles walks the template directory and runs template engine against passed data
func GenerateFiles(templatesFS embed.FS, templatePath string, targetRoot string, data interface{}) ([]string, error) {

	generatedFiles := []string{}
	err := fs.WalkDir(templatesFS, templatePath, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Create the same structure in the target directory
		relPath, _ := filepath.Rel(templatePath, path)
		targetPath := filepath.Join(targetRoot, filepath.Dir(relPath))
		fileName := filepath.Base(path)
		dst := filepath.Join(targetPath, fileName)
		err = os.MkdirAll(targetPath, os.ModePerm)

		if err != nil {
			return fmt.Errorf("error genering directory. %w", err)
		}

		// Handle template files
		// Copy or process other files as before
		if filepath.Ext(path) == ".mod" {
			dst = strings.Replace(dst, ".mod", "", 1)
		}

		if filepath.Ext(path) == ".tmpl" {
			dst = strings.Replace(dst, ".tmpl", ".go", 1)
		}

		if strings.HasPrefix(fileName, "hidden.") {
			dst = strings.Replace(dst, "hidden.", ".", 1)
			// just copy the files to the target directory
		}

		generatedFiles = append(generatedFiles, dst)

		return execTemplate(templatesFS, path, dst, data)
	})

	return generatedFiles, err
}

// execTemplate executes the template at destination directory
func execTemplate(f embed.FS, path, dst string, data interface{}) (err error) {
	tmpl, err := template.ParseFS(f, path)
	if err != nil {
		return fmt.Errorf("parsing template error. %w", err)
	}

	file, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("creating file error. path: %s  err: %w", dst, err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("closing file error. path: %s err: %w", dst, err)
		}
	}()

	err = tmpl.Execute(file, data)
	return
}
