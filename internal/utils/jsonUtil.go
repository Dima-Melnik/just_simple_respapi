package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		SendError(c, http.StatusBadRequest, err.Error())
		return false
	}

	return true
}
