package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/utils"
	"github.com/NewGlassbiller/go-sandbox/compiler/modulePath/gb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.SetOutput(os.Stderr)
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		var generateMethods, generateTypes, generateService bool
		outputPath := "./"
		options := plugin.Request.GetParameter()

		// Parse plugin options
		for _, option := range strings.Split(options, ",") {
			kv := strings.Split(option, "=")
			if len(kv) != 2 {
				continue
			}
			key, value := kv[0], kv[1]
			switch key {
			case "methods":
				generateMethods = (value == "true")
			case "types":
				generateTypes = (value == "true")
			case "services":
				generateService = (value == "true")
			case "out":
				outputPath = value
			}

		}

		if !generateMethods && !generateTypes && !generateService {
			fmt.Println("Please specify at least one of 'methods=true' or 'types=true'")
			os.Exit(1)
		}

		moduleData, gbPackageName := collectProtobufData(plugin)
		if err := generateTemplates(moduleData, outputPath); err != nil {
			return err
		}
		// fmt.Println("moduleData:", moduleData)
		fmt.Println("gbPackageName:", gbPackageName)
		if generateTypes {
			for modulePath, data := range moduleData {
				if err := generator.GenerateTypes(data.Messages, data.Enums, generateTypes, data.Services, modulePath, outputPath, HasTimestampFunc(data.Messages)); err != nil {
					return err
				}
			}
		}

		if generateMethods {
			for modulePath, data := range moduleData {
				if err := generator.GenerateMethods(data.Services, generateMethods, modulePath, outputPath); err != nil {
					return err
				}
			}
		}

		if generateService {
			for modulePath, data := range moduleData {
				fmt.Println("called")
				if err := generator.GenerateServices(data.Services, generateService, modulePath, outputPath); err != nil {
					return err
				}
				if err := generator.GenerateApp(data.Services, generateService, modulePath, outputPath); err != nil {
					return err
				}
				if err := generator.GenerateModule(data.Services, generateService, modulePath, outputPath); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func HasTimestampFunc(messages []generator.Message) bool {
	for _, message := range messages {
		for _, field := range message.Fields {
			if field.HasTimestamp {
				return true
			}
		}
	}
	return false
}

func collectProtobufData(plugin *protogen.Plugin) (map[string]struct {
	Messages []generator.Message
	Enums    []generator.Enum
	Services []generator.Service
}, string) {
	moduleData := make(map[string]struct {
		Messages []generator.Message
		Enums    []generator.Enum
		Services []generator.Service
	})
	var gbPackageName string

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		gbPackageName = *file.Proto.Package
		hasBgModulePath := proto.HasExtension(file.Desc.Options(), gb.E_GbModulePath)

		if hasBgModulePath {
			bgModulePath := proto.GetExtension(file.Desc.Options(), gb.E_GbModulePath)
			modulePath := bgModulePath.(string)

			// Initialize the entry if it doesn't exist
			if _, exists := moduleData[modulePath]; !exists {
				moduleData[modulePath] = struct {
					Messages []generator.Message
					Enums    []generator.Enum
					Services []generator.Service
				}{}
			}

			// Collect messages
			for _, msg := range file.Messages {
				var fields []generator.Field
				for _, field := range msg.Fields {
					var timestampVal = false

					if field.Desc.Kind().String() == "message" {
						if string(field.Desc.Message().FullName()) == "google.protobuf.Timestamp" {
							timestampVal = true
						}
					}
					fields = append(fields, generator.Field{
						Name:         string(field.Desc.Name()),
						Type:         utils.MapProtoType(field.Desc.Kind().String(), field),
						Number:       int(field.Desc.Number()),
						ProtoName:    string(field.Desc.Name()),
						HasTimestamp: timestampVal,
					})
				}

				// Get the current data for the module path
				moduleEntry := moduleData[modulePath]
				moduleEntry.Messages = append(moduleEntry.Messages, generator.Message{
					MessageName: string(msg.Desc.Name()),
					Fields:      fields,
				})

				// Update the map with the modified entry
				moduleData[modulePath] = moduleEntry
			}

			// Collect enums
			for _, enum := range file.Enums {
				var enumValues []generator.EnumValue
				for _, value := range enum.Values {
					enumValues = append(enumValues, generator.EnumValue{
						PName: string(enum.Desc.Name()),
						Name:  string(value.Desc.Name()),
						Value: int32(value.Desc.Number()),
					})
				}

				// Get the current data for the module path
				moduleEntry := moduleData[modulePath]
				moduleEntry.Enums = append(moduleEntry.Enums, generator.Enum{
					EnumName: string(enum.Desc.Name()),
					Values:   enumValues,
				})

				// Update the map with the modified entry
				moduleData[modulePath] = moduleEntry
			}

			// Collect services
			for _, service := range file.Services {
				var methods []generator.Method
				for _, method := range service.Methods {
					methods = append(methods, generator.Method{
						PackageName:  string(file.GoPackageName),
						PName:        string(service.Desc.Name()),
						MName:        string(method.Desc.Name()),
						RequestType:  method.Input.GoIdent.GoName,
						ResponseType: method.Output.GoIdent.GoName,
					})
				}

				// Get the current data for the module path
				moduleEntry := moduleData[modulePath]
				moduleEntry.Services = append(moduleEntry.Services, generator.Service{
					PackageName: string(file.GoPackageName),
					SName:       string(service.Desc.Name()),
					Methods:     methods,
				})

				// Update the map with the modified entry
				moduleData[modulePath] = moduleEntry
			}

		}
	}

	return moduleData, gbPackageName
}

func generateTemplates(moduleData map[string]struct {
	Messages []generator.Message
	Enums    []generator.Enum
	Services []generator.Service
}, outputPath string) error {
	for modulePath, data := range moduleData {
		// Extract the path after "internal/"
		internalSegment := "internal"
		index := strings.Index(modulePath, internalSegment)
		if index != -1 {
			relativePath := modulePath[index+len(internalSegment):]
			if strings.HasPrefix(relativePath, "/") {
				relativePath = strings.TrimPrefix(relativePath, "/")
			}
			modulePath = relativePath
		}

		// Create the directory structure
		moduleDir := filepath.Join(outputPath, "./templates", modulePath)
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
		for _, service := range data.Services {
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
