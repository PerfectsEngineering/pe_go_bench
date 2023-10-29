package formatter

import (
	"regexp"
)

// Matches all characters that are not in the valid syntax
var replacerRegex = regexp.MustCompile(`[^a-zA-Z0-9_]+?`)

func NormalizeWithRegex(text string) (string, error) {
	formattedString := replacerRegex.ReplaceAllString(text, "_")

	return formattedString, nil
}
