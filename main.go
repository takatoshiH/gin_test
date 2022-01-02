package main

import (
	"gin_test/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	albumEngine := engine.Group("/album")

	//一覧を取得する
	albumEngine.GET("/index", getAlbums)

	//新規に登録する
	albumEngine.POST("/create", controller.postAlbums)

	//IDからアルバムを取得する
	albumEngine.GET("/show/:id", controller.getAlbumByID)

	// 最も安いアルバムを取得する
	albumEngine.GET("/cheapest", controller.getCheapestAlbum)

	// localhost:8080
	engine.Run("localhost:8080")
}
