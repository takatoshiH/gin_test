package controller

import (
	"gin_test/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// HTMLを返すようにしたい
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// albumをIDは自動でインクリメントするようにしたい
func postAlbums(c *gin.Context) {
	var newAlbum model.Album

	//IDがincrementするようにする
	newAlbumID := getLatestID() + 1
	newAlbum.ID = strconv.Itoa(newAlbumID)

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	// jsonを返すようにしている
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	//こんな感じでパラメータを取得することができる
	id := c.Param("id")

	// Loopで一致するやつを探している
	for _, a := range albums {
		if a.ID == id {
			// jsonを返している(見つかった)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// jsonを返している(見つからなかった)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// 最も安いalbumを取得する
func getCheapestAlbum(c *gin.Context) {
	cheapestAlbum := albums[0]

	for _, a := range albums {
		if a.Price <= cheapestAlbum.Price {
			cheapestAlbum = a
		}
	}

	c.IndentedJSON(http.StatusOK, cheapestAlbum)
}

func getLatestID() int {
	index := 0

	for _, a := range albums {
		idInt, _ := strconv.Atoi(a.ID)
		if idInt > index {
			index = idInt
		}
	}

	return index
}
