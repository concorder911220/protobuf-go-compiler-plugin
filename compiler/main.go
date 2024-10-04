package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	log.SetOutput(os.Stderr)
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		var generateMethods, generateTypes, generateService bool
		options := plugin.Request.GetParameter()

		// Parse plugin options
		// Parse plugin options (e.g., "methods=true,types=false")
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
			}
		}

		if !generateMethods && !generateTypes && !generateService {
			fmt.Println("Please specify at least one of 'methods=true' or 'types=true'")
			os.Exit(1)
		}

		// Collect messages, enums, services
		messages, enums, services := collectProtobufData(plugin)

		// Generate types, methods, and service implementations
		if generateTypes {
			if err := generator.GenerateTypes(messages, enums, generateTypes, services, HasTimestampFunc(messages)); err != nil {
				return err
			}
		}

		if generateMethods {

			if err := generator.GenerateMethods(services, generateMethods); err != nil {
				return err
			}
		}

		if generateService {
			fmt.Println("services, ", services)
			if err := generator.GenerateServices(services, generateService); err != nil {

				return err
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

func collectProtobufData(plugin *protogen.Plugin) ([]generator.Message, []generator.Enum, []generator.Service) {
	var messages []generator.Message
	var enums []generator.Enum
	var services []generator.Service

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
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
			messages = append(messages, generator.Message{
				MessageName: string(msg.Desc.Name()),
				Fields:      fields,
			})
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
			enums = append(enums, generator.Enum{
				EnumName: string(enum.Desc.Name()),
				Values:   enumValues,
			})
		}

		// Collect services
		for _, service := range file.Services {
			var methods []generator.Method
			for _, method := range service.Methods {
				methods = append(methods, generator.Method{
					PName:        string(service.Desc.Name()),
					MName:        string(method.Desc.Name()),
					RequestType:  method.Input.GoIdent.GoName,
					ResponseType: method.Output.GoIdent.GoName,
				})
			}
			services = append(services, generator.Service{
				SName:   string(service.Desc.Name()),
				Methods: methods,
			})
		}
	}

	return messages, enums, services
}
