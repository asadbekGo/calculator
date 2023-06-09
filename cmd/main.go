package main

import (
	"fmt"

	"app/config"
	"app/controller"
	"app/models"
)

func main() {

	cfg := config.Load()

	con := controller.NewController(&cfg)

	users := con.UserGenerate(100)
	con.Users = users

	var (
		dataLimit int
		page      int
	)
	fmt.Println("Input Data limit:")
	fmt.Scan(&dataLimit)

	paginationCount := len(con.Users) / dataLimit
	fmt.Println("Pages count: ", paginationCount)

	for {
		fmt.Println("Input page:")
		fmt.Scan(&page)

		respUser := con.UserGetList(&models.UserGetListRequest{
			Offset: (page - 1) * dataLimit,
			Limit:  dataLimit,
		})

		for _, user := range respUser.Users {
			fmt.Println(user)
		}
	}
}
