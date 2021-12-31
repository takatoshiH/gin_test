package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)

	v1 := Book

	engine.Run(":3000")
}
