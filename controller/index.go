package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
