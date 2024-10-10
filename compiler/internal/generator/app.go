package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateServices generates the service implementations.
func GenerateApp(services []Service, generateService bool, modulePath string, outputPath string) error {
	if !generateService {
		return nil
	}
	// Prepare the data for the template
	data := struct {
		Services   []Service
		ModulePath string
	}{
		Services:   services,
		ModulePath: modulePath,
	}

	internalSegment := "internal"
	index := strings.Index(modulePath, internalSegment)
	if index != -1 {
		relativePath := modulePath[index+len(internalSegment):]
		if strings.HasPrefix(relativePath, "/") {
			relativePath = strings.TrimPrefix(relativePath, "/")
		}
		modulePath = relativePath
	}

	tmplPath := filepath.Join(outputPath, "templates", modulePath, "app", "app.tmpl")

	// Parse the template
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	// Execute the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// Format the generated code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting code: %v", err)
	}

	dirPath := filepath.Join(outputPath, "internal", modulePath, "app")

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory (with any necessary parent directories)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	// Write the generated service file
	_outputPath := filepath.Join(dirPath, "app.go")
	if err := os.WriteFile(_outputPath, formattedCode, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Generated: %s\n", _outputPath)
	return nil
}
