package helpers

import (
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/utils"
	moduleName "github.com/NewGlassbiller/go-sandbox/compiler/moduleName/gb"
	"github.com/NewGlassbiller/go-sandbox/compiler/modulePath/gb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func CollectProtobufData(plugin *protogen.Plugin, outputPath string) (map[string]struct {
	SupplyData generator.SupplyData
}, string) {
	moduleData := make(map[string]struct {
		SupplyData generator.SupplyData
	})
	var gbPackageName string

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		gbPackageName = *file.Proto.Package
		hasBgModulePath := proto.HasExtension(file.Desc.Options(), gb.E_GbModulePath)
		hasBgModuleName := proto.HasExtension(file.Desc.Options(), moduleName.E_GbModuleName)

		if hasBgModulePath && hasBgModuleName {
			bgModuleName := proto.GetExtension(file.Desc.Options(), moduleName.E_GbModuleName)
			moduleName := bgModuleName.(string)
			bgModulePath := proto.GetExtension(file.Desc.Options(), gb.E_GbModulePath)
			modulePath := bgModulePath.(string)

			// Initialize the entry if it doesn't exist
			if _, exists := moduleData[moduleName]; !exists {
				moduleData[moduleName] = struct {
					SupplyData generator.SupplyData
				}{}
			}

			moduleEntry := moduleData[moduleName]
			moduleEntry.SupplyData.MetaInfo = generator.MetaInfo{
				ModuleName: moduleName,
				ModulePath: modulePath,
				OutputPath: outputPath,
			}

			// Update the map with the modified entry
			moduleData[moduleName] = moduleEntry

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
				moduleEntry := moduleData[moduleName]
				moduleEntry.SupplyData.TypeData.Messages = append(moduleEntry.SupplyData.TypeData.Messages, generator.Message{
					MessageName: string(msg.Desc.Name()),
					Fields:      fields,
				})

				// Update the map with the modified entry
				moduleData[moduleName] = moduleEntry
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
				moduleEntry := moduleData[moduleName]
				moduleEntry.SupplyData.TypeData.Enums = append(moduleEntry.SupplyData.TypeData.Enums, generator.Enum{
					EnumName: string(enum.Desc.Name()),
					Values:   enumValues,
				})

				// Update the map with the modified entry
				moduleData[moduleName] = moduleEntry
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
				moduleEntry := moduleData[moduleName]
				moduleEntry.SupplyData.TypeData.Services = append(moduleEntry.SupplyData.TypeData.Services, generator.Service{
					PackageName: string(file.GoPackageName),
					SName:       string(service.Desc.Name()),
					Methods:     methods,
				})

				// Update the map with the modified entry
				moduleData[moduleName] = moduleEntry
			}

		}
	}

	return moduleData, gbPackageName
}
