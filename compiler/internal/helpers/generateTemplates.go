package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
)

func GenerateTemplates(moduleData map[string]struct {
	SupplyData generator.SupplyData
}, outputPath string, templatePath string) error {
	for _, data := range moduleData {
		moduleDir := filepath.Join(outputPath, templatePath, data.SupplyData.MetaInfo.ModuleName)
		if err := os.MkdirAll(moduleDir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %s: %v", moduleDir, err)
		}

		// Generate module.tmpl
		moduleFilePath := filepath.Join(moduleDir, "module.tmpl")

		if _, err := os.Stat(moduleFilePath); os.IsNotExist(err) {
			defaultAppTemplate := `package app

	import (
		"github.com/NewGlassbiller/gb-go-common/gbnet/common"
		"github.com/NewGlassbiller/gb-go-common/gbpolicy"
		"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/dto"
		"google.golang.org/grpc/codes"
		"google.golang.org/grpc/status"
		"gorm.io/gorm"
	)

	
	`

			// Ensure the "internal/templates" directory exists
			if err := os.MkdirAll(filepath.Dir(moduleFilePath), os.ModePerm); err != nil {
				log.Fatalf("Failed to create directories: %v", err)
			}

			// Write the default template content to the file
			err = os.WriteFile(moduleFilePath, []byte(defaultAppTemplate), 0644)
			if err != nil {
				log.Fatalf("Failed to create template file: %v", err)
			}

			log.Println("Template file created:", moduleFilePath)

		} else {
			log.Println("Template file already exists:", moduleFilePath)
		}

		// Generate service_gen.tmpl
		serviceFilePath := filepath.Join(moduleDir, "service_gen.tmpl")
		if _, err := os.Stat(serviceFilePath); os.IsNotExist(err) {
			defaultAppTemplate := `package app

	import (
		"github.com/NewGlassbiller/gb-go-common/gbnet/common"
		"github.com/NewGlassbiller/gb-go-common/gbpolicy"
		"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/dto"
		"google.golang.org/grpc/codes"
		"google.golang.org/grpc/status"
		"gorm.io/gorm"
	)

	
	`

			// Ensure the "internal/templates" directory exists
			if err := os.MkdirAll(filepath.Dir(serviceFilePath), os.ModePerm); err != nil {
				log.Fatalf("Failed to create directories: %v", err)
			}

			// Write the default template content to the file
			err = os.WriteFile(serviceFilePath, []byte(defaultAppTemplate), 0644)
			if err != nil {
				log.Fatalf("Failed to create template file: %v", err)
			}

			log.Println("Template file created:", serviceFilePath)

		} else {
			log.Println("Template file already exists:", serviceFilePath)
		}

		typesFilePath := filepath.Join(moduleDir, "type.tmpl")
		if _, err := os.Stat(typesFilePath); os.IsNotExist(err) {
			defaultAppTemplate := `package app

	import (
		"github.com/NewGlassbiller/gb-go-common/gbnet/common"
		"github.com/NewGlassbiller/gb-go-common/gbpolicy"
		"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/dto"
		"google.golang.org/grpc/codes"
		"google.golang.org/grpc/status"
		"gorm.io/gorm"
	)


	`

			// Ensure the "internal/templates" directory exists
			if err := os.MkdirAll(filepath.Dir(typesFilePath), os.ModePerm); err != nil {
				log.Fatalf("Failed to create directories: %v", err)
			}

			// Write the default template content to the file
			err = os.WriteFile(typesFilePath, []byte(defaultAppTemplate), 0644)
			if err != nil {
				log.Fatalf("Failed to create template file: %v", err)
			}

			log.Println("Template file created:", typesFilePath)

		} else {
			log.Println("Template file already exists:", typesFilePath)
		}

		// Create app directory
		appDir := filepath.Join(moduleDir, "app")
		if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating app directory %s: %v", appDir, err)
		}

		// Generate app.tmpl
		appFilePath := filepath.Join(appDir, "app.tmpl")
		if _, err := os.Stat(appFilePath); os.IsNotExist(err) {
			defaultAppTemplate := `package app

	import (
		"github.com/NewGlassbiller/gb-go-common/gbnet/common"
		"github.com/NewGlassbiller/gb-go-common/gbpolicy"
		"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/dto"
		"google.golang.org/grpc/codes"
		"google.golang.org/grpc/status"
		"gorm.io/gorm"
	)

	
	`

			// Ensure the "internal/templates" directory exists
			if err := os.MkdirAll(filepath.Dir(appFilePath), os.ModePerm); err != nil {
				log.Fatalf("Failed to create directories: %v", err)
			}

			// Write the default template content to the file
			err = os.WriteFile(appFilePath, []byte(defaultAppTemplate), 0644)
			if err != nil {
				log.Fatalf("Failed to create template file: %v", err)
			}

			log.Println("Template file created:", appFilePath)

		} else {
			log.Println("Template file already exists:", appFilePath)
		}

		// Generate method templates for each service
		for _, service := range data.SupplyData.TypeData.Services {
			for _, method := range service.Methods {
				methodFilePath := filepath.Join(appDir, fmt.Sprintf("%s.tmpl", method.MName))

				// Check if the template file exists
				if _, err := os.Stat(methodFilePath); os.IsNotExist(err) {
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
					if err := os.MkdirAll(filepath.Dir(methodFilePath), os.ModePerm); err != nil {
						log.Fatalf("Failed to create directories: %v", err)
					}

					// Write the default template content to the file
					err = os.WriteFile(methodFilePath, []byte(defaultTemplate), 0644)
					if err != nil {
						log.Fatalf("Failed to create template file: %v", err)
					}

					log.Println("Template file created:", methodFilePath)
				} else {
					log.Println("Template file already exists:", methodFilePath)
				}

			}
		}
	}
	return nil
}
