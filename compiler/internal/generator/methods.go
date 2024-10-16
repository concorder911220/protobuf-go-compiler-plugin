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

	"github.com/NewGlassbiller/go-sandbox/compiler/internal/utils"
)

// GenerateMethods generates the method handler templates.
func GenerateMethods(SupplyData SupplyData, outputPath string, templatePath string) error {

	for _, service := range SupplyData.TypeData.Services {
		for _, method := range service.Methods {
			data := struct {
				ServiceName      string
				ServiceNameLower string
				MethodName       string
				RequestType      string
				ResponseType     string
				PackageName      string
				ModuleName       string
				ModulePath       string
				IsCommand        bool
			}{
				ServiceName:      service.SName,
				ServiceNameLower: strings.ToLower(service.SName),
				MethodName:       method.MName,
				RequestType:      method.RequestType,
				ResponseType:     method.ResponseType,
				PackageName:      service.PackageName,
				ModuleName:       SupplyData.MetaInfo.ModuleName,
				ModulePath:       SupplyData.MetaInfo.ModulePath,
				IsCommand:        method.IsCommand,
			}

			prefix := "q_" // Default prefix
			if method.IsCommand {
				prefix = "cmd_" // Change prefix if IsCommand is true
			}

			tmplPath := filepath.Join(templatePath, data.ModuleName, "app", fmt.Sprintf("%s_%s.tmpl", prefix, method.MName))
			fmt.Println("tmplPath:", tmplPath)

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

			dirPath := filepath.Join(outputPath, data.ModuleName, "app")

			// Check if the directory exists
			if _, err := os.Stat(dirPath); os.IsNotExist(err) {
				// Create the directory (with any necessary parent directories)
				err := os.MkdirAll(dirPath, os.ModePerm)
				if err != nil {
					log.Fatalf("Failed to create directory: %v", err)
				}
			}
			// Prepare output path
			fileName := utils.MethodToSnakeCase(method.MName)
			outputPath := filepath.Join(dirPath, fileName)

			// Check if file exists and prompt for overwrite
			if err := utils.WriteFileWithPrompt(outputPath, formattedCode); err != nil {
				return err
			}

			fmt.Printf("Generated: %s\n", outputPath)
		}
	}

	return nil
}
