package main

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/bdtomlin/gostak/web/component"
)

func main() {
	defer model.Init()()

	// fmt.Println(component.CssClass("one one", "two two"))
	// fmt.Println(component.CssClassOverride("one one", "two two"))
	// fmt.Println(component.CssClassOverride("one one", "two two", "three three"))
	// fmt.Println(component.CssClassOverride("one one", " "))
}
