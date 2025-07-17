package handlers

import (
	"just_simple_api/internal/auth"
	"just_simple_api/internal/models"
	"just_simple_api/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandlers struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandlers {
	return &AuthHandlers{db: db}
}

func (h *AuthHandlers) Registere(c *gin.Context) {
	var user models.ShopUser
	if !utils.BindJSON(c, &user) {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	hashedPasswoord, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = hashedPasswoord

	if err := h.db.Create(&user).Error; err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (h *AuthHandlers) Login(c *gin.Context) {
	var user models.ShopUser
	var input models.ShopUser

	if !utils.BindJSON(c, &input) {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.db.Where("name ILIKE ?", input.Name).First(&user).Error; err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.SendError(c, http.StatusInternalServerError, "user not found")
		return
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
