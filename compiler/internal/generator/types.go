package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateTypes generates Go types and interfaces from protobuf messages.
func GenerateTypes(messages []Message, enums []Enum, generateTypes bool, services []Service, hasTimestamp bool) error {
	if !generateTypes {
		return nil
	}
	tmplPath := filepath.Join(".", "templates", "types.tmpl")

	// Check if the template file exists
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		// Create the template file with the default content
		defaultTemplate := `package app

{{ if .HasTimestamp }}
import timestamppb "google.golang.org/protobuf/types/known/timestamppb"
{{ end }}

{{- range .Enums }}
type {{ .EnumName }} int32

const (
{{- range .Values }}
	{{ .Name }} {{ .PName }} = {{ .Value }} 
{{- end }}
)
{{- end }}

{{- range .Messages }}
type {{ .MessageName }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }}
{{- end }}
}
{{- end }}

{{- range .Services }}
type {{ .SName }} interface {
	{{- range .Methods }}
	{{ .MName }}(req *{{ .RequestType }}) (*{{ .ResponseType }}, error)
	{{- end }}
}
{{- end }}`

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

	// Parse the template
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	// Data to be passed to the template
	data := struct {
		Messages     []Message
		Enums        []Enum
		Services     []Service
		HasTimestamp bool
	}{
		Messages:     messages,
		Enums:        enums,
		Services:     services,
		HasTimestamp: hasTimestamp,
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

	dirPath := filepath.Join(".", "interface")

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory (with any necessary parent directories)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	// Write the generated file
	outputPath := filepath.Join(dirPath, "types_and_interfaces.go")
	if err := os.WriteFile(outputPath, formattedCode, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Generated: %s\n", outputPath)
	return nil
}
