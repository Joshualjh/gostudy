package main

import (
	"fmt"
	"main/pkg/api"
	"main/pkg/db"
	"main/pkg/router"
)

func main() {

	dbHandler, err := db.NewAndConnectGorm("username:password@tcp(url)/dev?charset=utf8mb4&parseTime=True&loc=Local&tls=true")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8081")
}
