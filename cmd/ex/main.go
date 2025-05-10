package main

import (
	"fmt"
	"net/mail"

	"github.com/bdtomlin/gostak/internal/repo"
)

func main() {
	repo.Init()
	add, err := mail.ParseAddress("bob bob<bob@bob.com>")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(add.Name)
	fmt.Println(add.Address)

	add, err = mail.ParseAddress("bobb")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(add.Name)
		fmt.Println(add.Address)
	}
}
