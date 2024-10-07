package utils

import (
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
)

// MethodToSnakeCase converts method names to snake_case.
func MethodToSnakeCase(methodName string) string {
	var output []rune
	for i, r := range methodName {
		if unicode.IsUpper(r) && i > 0 {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output) + ".go"
}

// MapProtoType maps protobuf types to Go types.
func MapProtoType(protoType string, field *protogen.Field) string {
	var goType string
	switch protoType {
	case "string":
		goType = "string"
	case "int32":
		goType = "int32"
	case "bool":
		goType = "bool"
	case "float":
		goType = "float64"
	case "message":
		{
			if string(field.Desc.Message().FullName()) == "google.protobuf.Timestamp" {
				goType = "*timestamppb.Timestamp"
			} else {
				goType = "*" + string(field.Desc.Message().Name()) // Use pointer to the message type
			}

		}

	case "enum":
		goType = string(field.Desc.Enum().Name()) // Use pointer to the enum type
	// Add more mappings as necessary
	default:
		goType = "interface{}" // Fallback for unknown types
	}

	if field.Desc.Cardinality() == 3 {
		return "[]" + goType // Return slice for repeated fields
	} else if field.Desc.HasOptionalKeyword() {
		return "*" + goType // Return pointer for optional fields
	}

	// Return the basic type for normal fields
	return goType
}

// ToCamelCase converts snake_case to CamelCase.
func ToCamelCase(snake string) string {
	parts := strings.Split(snake, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part) // Capitalize each part
	}
	return strings.Join(parts, "")
}
