package main

import (
	"fmt"

	"app/controller"
)

func main() {

	users := controller.GenerateUser(10)
	for _, user := range users {
		fmt.Println(user)
	}
}
