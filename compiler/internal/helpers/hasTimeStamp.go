package helpers

import (
	"github.com/NewGlassbiller/go-sandbox/compiler/internal/generator"
)

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
