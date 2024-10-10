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
func GenerateMethods(services []Service, generateMethods bool, modulePath string, outputPath string) error {
	if !generateMethods {
		return nil
	}

	for _, service := range services {
		for _, method := range service.Methods {
			data := struct {
				ServiceName      string
				ServiceNameLower string
				MethodName       string
				RequestType      string
				ResponseType     string
				PackageName      string
				ModulePath       string
			}{
				ServiceName:      service.SName,
				ServiceNameLower: strings.ToLower(service.SName),
				MethodName:       method.MName,
				RequestType:      method.RequestType,
				ResponseType:     method.ResponseType,
				PackageName:      service.PackageName,
				ModulePath:       modulePath,
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

			tmplPath := filepath.Join(outputPath, "templates", modulePath, "app", fmt.Sprintf("%s.tmpl", method.MName))
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

			dirPath := filepath.Join(outputPath, "internal", modulePath, "app")

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
