package main

import (
	"fmt"

	"github.com/bdtomlin/gostak/internal/repo"
)

func main() {
	repo.Init()
	// users, _ := repo.ListUsers()
	// fmt.Println(users)

	user, err := repo.GetUser("01961d19-5aa9-75d7-946f-77457a8d67e8")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", user)
		fmt.Println(user.ID)
	}
}
