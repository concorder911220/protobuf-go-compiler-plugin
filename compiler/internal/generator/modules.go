package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateModules generates all go files from templates.
func GenerateModules(SupplyData SupplyData, outputPath string, templatePath string, hasTimestamp bool) error {
	data := struct {
		Messages     []Message
		Enums        []Enum
		Services     []Service
		HasTimestamp bool
		ModulePath   string
		ModuleName   string
	}{
		Messages:     SupplyData.TypeData.Messages,
		Enums:        SupplyData.TypeData.Enums,
		Services:     SupplyData.TypeData.Services,
		HasTimestamp: hasTimestamp,
		ModulePath:   SupplyData.MetaInfo.ModulePath,
		ModuleName:   SupplyData.MetaInfo.ModuleName,
	}

	// Define the templates directory
	templatesDir := filepath.Join(templatePath, data.ModuleName)
	fmt.Println("Templates Directory:", templatesDir)

	// List only the root-level template files
	files, err := os.ReadDir(templatesDir)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", templatesDir, err)
	}

	// Process only the .tmpl files in the root directory and app/app.tmpl
	for _, file := range files {
		if file.IsDir() {
			// Check for the "app" subdirectory
			if file.Name() == "app" {
				// Read the app/app.tmpl file
				appTemplatePath := filepath.Join(templatesDir, "app", "app.tmpl")
				if err := processTemplate(appTemplatePath, data, outputPath, data.ModuleName, "app"); err != nil {
					return err
				}
			}
			continue // Skip other subdirectories
		}

		// Check if the file is a .tmpl file in the root directory
		if strings.HasSuffix(file.Name(), ".tmpl") {
			templatePath := filepath.Join(templatesDir, file.Name())
			if err := processTemplate(templatePath, data, outputPath, data.ModuleName, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

// processTemplate handles parsing, executing, formatting, and writing the template file.
func processTemplate(templatePath string, data interface{}, outputPath, modulePath, subDir string) error {
	// Parse the template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing template %s: %v", templatePath, err)
	}

	// Execute the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("error executing template %s: %v", templatePath, err)
	}

	// Format the generated code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting code: %v", err)
	}

	// Prepare output directory based on the subdirectory
	dirPath := filepath.Join(outputPath, modulePath, subDir)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory (with any necessary parent directories)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Write the generated Go file
	outputFileName := strings.TrimSuffix(filepath.Base(templatePath), ".tmpl") + ".go"
	outputFilePath := filepath.Join(dirPath, outputFileName)
	fmt.Println("Output Path:", outputFilePath)
	if err := os.WriteFile(outputFilePath, formattedCode, 0644); err != nil {
		return fmt.Errorf("error writing file %s: %v", outputFilePath, err)
	}

	fmt.Printf("Generated: %s\n", outputFilePath)
	return nil
}
