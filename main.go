package main

import (
	"gin_test/server"
)

func main() {
	router := server.GetRouter()

	// DBに繋いだりする
	//db.Init()
	router.Run(":8080")
}
