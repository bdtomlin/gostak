package main

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/bdtomlin/gostak/web/component"
)

func main() {
	model.Init()
	defer model.Close()

	fmt.Println(component.CssClass("one one", "two two"))
	fmt.Println(component.CssClassOverride("one one", "two two"))
	fmt.Println(component.CssClassOverride("one one", "two two", "three three"))
	fmt.Println(component.CssClassOverride("one one", " "))
}
