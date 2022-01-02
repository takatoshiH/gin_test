package server

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	//一覧を取得する
	router.GET("/albums", getAlbums)

	//新規に登録する
	router.POST("/albums", postAlbums)

	//IDからアルバムを取得する
	router.GET("/albums/:id", getAlbumByID)

	// 最も安いアルバムを取得する
	router.GET("/albums/cheapest", getCheapestAlbum)

	return router
}
