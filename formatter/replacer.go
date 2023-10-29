package formatter

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

var validStringRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var replacer = strings.NewReplacer(
	" ", "_",
	",", "_",
	".", "_",
	"-", "_",
	"*", "_",
	"!", "_",
	"?", "_",
	"(", "_",
	")", "_",
)

func NormalizeWithReplacer(text string) (string, error) {
	formattedString := replacer.Replace(text)

	if !validStringRegex.MatchString(formattedString) {
		slog.Error("error normalising string", "value", text, "normalizedValue", formattedString)
		return "", fmt.Errorf("error normalising string %s", text)
	}

	return formattedString, nil
}
