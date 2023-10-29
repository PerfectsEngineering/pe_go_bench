package formatter_test

import (
	"testing"

	"github.com/perfectsengineering/pe_go_bench/formatter"
)

func BenchmarkNormalise(b *testing.B) {
	b.Run("WithReplacer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			formatter.NormalizeWithReplacer("Hello, World!")
		}
	})

	b.Run("WithRegexp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			formatter.NormalizeWithRegex("Hello, World!")
		}
	})
}

func TestNormaliseSuccess(t *testing.T) {
	testcases := []struct {
		text           string
		normalizedText string
	}{
		{
			text:           "Hello, World!",
			normalizedText: "Hello__World_",
		},
		{
			text:           "name.text.based",
			normalizedText: "name_text_based",
		},
		{
			text:           "name-text_multiple*",
			normalizedText: "name_text_multiple_",
		},
	}

	t.Run("WithReplacer", func(t *testing.T) {
		for _, tc := range testcases {
			normalizedText, err := formatter.NormalizeWithReplacer(tc.text)
			if err != nil {
				t.Fatal(err)
			}

			if normalizedText != tc.normalizedText {
				t.Fatalf("expected %s got %s", tc.normalizedText, normalizedText)
			}
		}
	})

	t.Run("WithRegexp", func(t *testing.T) {
		for _, tc := range testcases {
			normalizedText, err := formatter.NormalizeWithRegex(tc.text)
			if err != nil {
				t.Fatal(err)
			}

			if normalizedText != tc.normalizedText {
				t.Fatalf("expected %s got %s", tc.normalizedText, normalizedText)
			}
		}
	})
}

// Other tests cases can be written for failure cases etc
