package component

import (
	"testing"

	"github.com/bdtomlin/gostak/web/component"
)

func TestCssClass(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"one one", "two two"}, "one one two two"},
		{[]string{"one one", ""}, "one one"},
	}

	for _, test := range tests {
		result := component.CssClass(test.input...)
		if result != test.expected {
			t.Errorf("CssClass(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestCssClassOverride(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"one one", "two two"}, "two two"},
		{[]string{"one one", ""}, "one one"},
		{[]string{"one one", "", ""}, "one one"},
	}

	for _, test := range tests {
		result := component.CssClassOverride(test.input...)
		if result != test.expected {
			t.Errorf("CssClass(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
