package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/helpers"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	log.SetOutput(os.Stderr)
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		var generateMethods, generateModules bool
		outputPath := "./"
		templatePath := "./"
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
			case "modules":
				generateModules = (value == "true")
			case "out":
				outputPath = value
			case "tplpath":
				templatePath = value
			}

		}

		if !generateMethods && !generateModules {
			fmt.Println("Please specify at least one of 'methods=true' or 'modules=true'")
			os.Exit(1)
		}

		moduleData, gbPackageName := helpers.CollectProtobufData(plugin, outputPath)
		fmt.Println("gbPackageName:", gbPackageName)
		if generateModules {
			for _, data := range moduleData {
				if err := generator.GenerateModules(data.SupplyData, outputPath, templatePath, helpers.HasTimestampFunc(data.SupplyData.TypeData.Messages)); err != nil {
					return err
				}
			}
		}

		if generateMethods {
			for _, data := range moduleData {
				if err := generator.GenerateMethods(data.SupplyData, outputPath, templatePath); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
