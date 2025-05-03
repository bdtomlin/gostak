package main

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/web/form"
)

func main() {
	repo.Init()
	su := form.NewSignUp()
	su.Email = "onetwo.three"
	su.Password = "1234"
	su.Validate()
	fmt.Println("su:", su)
	fmt.Printf("form: %+v\n", su.Form)
	fmt.Println("errors:", su.Errors)
	fmt.Println("form:", su.Form)
}
