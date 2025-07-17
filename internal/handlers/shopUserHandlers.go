package handlers

import (
	"just_simple_api/internal/handlers/services"
	"just_simple_api/internal/models"
	"just_simple_api/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShopUserHandlers struct {
	service services.UserShopServices
}

func NewShopUserHandlers(s services.UserShopServices) *ShopUserHandlers {
	return &ShopUserHandlers{service: s}
}

func (h ShopUserHandlers) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 10 {
		limit = 10
	}

	offSet := (page - 1) * limit

	users, err := h.service.GetAllUsers(offSet, limit)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h ShopUserHandlers) GetUserByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h ShopUserHandlers) AddUser(c *gin.Context) {
	var createdUser models.ShopUser

	if !utils.BindJSON(c, &createdUser) {
		return
	}

	if err := h.service.AddUser(createdUser); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h ShopUserHandlers) UpdateUser(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedUser models.ShopUser
	if !utils.BindJSON(c, &updatedUser) {
		return
	}

	if err := h.service.UpdateUser(updatedUser, id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h ShopUserHandlers) DeleteUser(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Deleted user")
}
