package utils

import (
	"fmt"
	"os"
	"strings"
)

// WriteFileWithPrompt writes the formatted code to a file, asking for confirmation if the file already exists.
func WriteFileWithPrompt(filePath string, formattedCode []byte) error {
	if _, err := os.Stat(filePath); err == nil {
		if !promptOverwrite(filePath) {
			fmt.Printf("Skipped: %s\n", filePath)
			return nil
		}
	}
	return os.WriteFile(filePath, formattedCode, 0644)
}

// promptOverwrite prompts the user to confirm if an existing file should be overwritten.
func promptOverwrite(filePath string) bool {
	var response string
	fmt.Printf("File %s already exists. Overwrite? (y/n): ", filePath)
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}
