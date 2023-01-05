package main

import (
	"fmt"

	"main.go/pkg/api"
	"main.go/pkg/db"
	"main.go/pkg/router"
)

func main() {
	dbHandler, err := db.NewConnectGorm("cruduser:sADMIN123@tcp(crudserverset.mysql.database.azure.com)/dev?charset=utf8mb4&parseTime=True&loc=Local&tls=true")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8082")
}
