package main

import "main/pkg/router"

func main() {
	r := router.Router()

	r.Run(":8081")
}
