package htm

import (
	"testing"
)

func TestAssetPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "image.jpg",
			expected: "image-thehashhere.jpg",
		},
		{
			input:    "script.min.js",
			expected: "script-thehashhere.min.js",
		},
		{
			input:    "styles.css",
			expected: "styles-thehashhere.css",
		},
		{
			input:    "path.to.file.txt",
			expected: "path-thehashhere.to.file.txt",
		},
		{
			input:    "singlepart.tar.gz",
			expected: "singlepart-thehashhere.tar.gz",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := AssetPath(test.input)
			if result != test.expected {
				t.Errorf("AssetPath(%q) = %q, expected %q", test.input, result, test.expected)
			}
		})
	}
}
