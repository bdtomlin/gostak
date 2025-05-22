package main

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/web/component"
)

func main() {
	repo.Init()

	fmt.Println(component.CssClass("one one", "two two"))
	fmt.Println(component.CssClassOverride("one one", "two two"))
	fmt.Println(component.CssClassOverride("one one", "two two", "three three"))
	fmt.Println(component.CssClassOverride("one one", " "))
}
