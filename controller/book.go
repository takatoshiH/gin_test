package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiita/model"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
}
