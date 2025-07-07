package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckCorrectID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		SendError(c, http.StatusBadRequest, "Invalid ID")
		return id, false
	}

	return id, true
}
