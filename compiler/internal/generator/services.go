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

// GenerateServices generates the service implementations.
func GenerateServices(services []Service, generateService bool) error {
	if !generateService {
		return nil
	}

	tmplPath := filepath.Join(".", "templates", "service.tmpl")

	// Check if the template file exists
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		// Create the template file with the default content
		defaultTemplate := `package claim

import (
	"context"

	"github.com/NewGlassbiller/gb-go-common/gbnet/common"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/app"
	dto "github.com/NewGlassbiller/gb-services-insurance/pkg/proto/claim/dto"
	grpcgen "github.com/NewGlassbiller/gb-services-insurance/pkg/proto/claim/grpc"
)

{{- range .Services }}
type {{ .SName }}Service struct {
	grpcgen.Unimplemented{{ .SName }}Server
	command *app.Command
	query   *app.Query
}

func New{{ .SName }}Service(command *app.Command, query *app.Query) *{{ .SName }}Service {
	return &{{ .SName }}Service{
		command: command,
		query:   query,
	}
}

{{- range .Methods }}
func (s *{{ .PName }}Service) {{ .MName }}(ctx context.Context, req *dto.{{ .RequestType }}) (*dto.{{ .ResponseType }}, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.{{ .MName }}(&authUser, req)
}
{{- end }}
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

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	// Prepare the data for the template
	data := struct {
		Services []Service
	}{
		Services: services,
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

	dirPath := filepath.Join(".", "internal")

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory (with any necessary parent directories)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	// Write the generated service file
	outputPath := filepath.Join(dirPath, "service.go")
	if err := os.WriteFile(outputPath, formattedCode, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Generated: %s\n", outputPath)
	return nil
}
