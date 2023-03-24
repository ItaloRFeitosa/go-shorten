package main

import "github.com/italorfeitosa/go-shorten/internal/linkmanager"

func main() {
	di := linkmanager.Setup()

	defer func(di *linkmanager.Container) {
		db, _ := di.DB.DB()
		db.Close()
	}(di)

	di.FiberApp.Listen(":8080")
}
