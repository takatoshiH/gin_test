package server

import (
	"gin_test/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	album := new(controller.AlbumController)

	//一覧を取得する
	router.GET("/", album.Index)

	//新規に登録する
	router.POST("/albums", album.Create)

	//IDからアルバムを取得する
	router.GET("/albums/:id", album.Show)

	// 最も安いアルバムを取得する
	router.GET("/albums/cheapest", album.GetCheapestAlbum)

	return router
}
