package main

import (
	"log"
	"terraform-provider-example/api/server"
)

func main() {
	items := map[string]server.Item{}
	itemService := server.NewService(items)
	err := itemService.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
