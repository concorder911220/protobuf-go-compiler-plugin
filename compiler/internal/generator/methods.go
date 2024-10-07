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
func GenerateMethods(services []Service, generateMethods bool) error {
	if !generateMethods {
		return nil
	}

	tmplPath := filepath.Join(".", "templates", "command.tmpl")

	// Check if the template file exists
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		// Create the template file with the default content
		defaultTemplate := `package app

import (
	"github.com/NewGlassbiller/gb-go-common/gbnet/common"
	"github.com/NewGlassbiller/gb-go-common/gbpolicy"
	"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/{{ .ServiceNameLower }}/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// {{ .MethodName }} handles {{ .RequestType }} and returns {{ .ResponseType }}
func (cmd Command) {{ .MethodName }}(authUser *common.UserInfo, req *dto.{{ .RequestType }}) (*dto.{{ .ResponseType }}, error) {
	// TODO: Add your logic here for {{ .MethodName }}
	return nil, nil
}`

		// Ensure the "internal/templates" directory exists
		if err := os.MkdirAll(filepath.Dir(tmplPath), os.ModePerm); err != nil {
			log.Fatalf("Failed to create directories: %v", err)
		}

		// Write the default template content to the file
		err = os.WriteFile(tmplPath, []byte(defaultTemplate), 0644)
		if err != nil {
			log.Fatalf("Failed to create template file: %v", err)
		}

		log.Println("Template file created:", tmplPath)
	} else {
		log.Println("Template file already exists:", tmplPath)
	}

	for _, service := range services {
		for _, method := range service.Methods {
			data := struct {
				ServiceName      string
				ServiceNameLower string
				MethodName       string
				RequestType      string
				ResponseType     string
			}{
				ServiceName:      service.SName,
				ServiceNameLower: strings.ToLower(service.SName),
				MethodName:       method.MName,
				RequestType:      method.RequestType,
				ResponseType:     method.ResponseType,
			}

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

			dirPath := filepath.Join(".", "internal/app")

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
