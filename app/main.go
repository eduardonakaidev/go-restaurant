package main

import (
	"github.com/eduardonakaidev/go-restaurant/connection"
	"github.com/eduardonakaidev/go-restaurant/router"
	"github.com/eduardonakaidev/go-restaurant/utils"
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
