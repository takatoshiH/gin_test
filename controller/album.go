package controller

import (
	. "gin_test/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumController struct{}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (a AlbumController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"albums": albums,
	})
}

func (a AlbumController) Create(c *gin.Context) {
	var newAlbum Album

	newAlbumID := getLatestID() + 1
	newAlbum.ID = strconv.Itoa(newAlbumID)

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	// jsonを返すようにしている
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (a AlbumController) Show(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (a AlbumController) GetCheapestAlbum(c *gin.Context) {
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
