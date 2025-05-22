package component

import (
	"slices"
	"strings"
)

func CssClass(classes ...string) string {
	classes = slices.DeleteFunc(classes, func(cls string) bool {
		return strings.TrimSpace(cls) == ""
	})
	return strings.Join(classes, " ")
}

func CssClassOverride(classes ...string) string {
	classes = slices.DeleteFunc(classes, func(cls string) bool {
		return strings.TrimSpace(cls) == ""
	})

	if len(classes) > 1 {
		classes = classes[1:]
	}
	return strings.Join(classes, " ")
}
