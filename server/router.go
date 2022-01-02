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
	router.GET("/", album.GetAlbums)

	//新規に登録する
	router.POST("/albums", album.PostAlbums)

	//IDからアルバムを取得する
	router.GET("/albums/:id", album.GetAlbumByID)

	// 最も安いアルバムを取得する
	router.GET("/albums/cheapest", album.GetCheapestAlbum)

	return router
}
