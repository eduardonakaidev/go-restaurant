package main

import (
	"github.com/eduardonakaidev/go-restaurant/app/connection"
	"github.com/eduardonakaidev/go-restaurant/app/router"
	"github.com/eduardonakaidev/go-restaurant/app/utils"
)

func init() {
	err := utils.LoadEnv()
	if err != nil {
		panic(err)
	}
}

func main() {
	connection.ConnectPostgresDB()
	router.HandleRequest()
}