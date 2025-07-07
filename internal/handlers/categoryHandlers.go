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

type CategoryHandlers struct {
	service services.CategoryServices
}

func NewCategoryHandlers(s services.CategoryServices) *CategoryHandlers {
	return &CategoryHandlers{service: s}
}

func (h CategoryHandlers) GetAllCategory(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page > 1 {
		page = 1
	}

	if limit > 1 {
		limit = 10
	}

	offSet := (page - 1) * limit

	categories, err := h.service.GetAllCategory(offSet, limit)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h CategoryHandlers) GetCategoryByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h CategoryHandlers) CreateCategory(c *gin.Context) {
	var createdCategory models.Category
	if !utils.BindJSON(c, createdCategory) {
		return
	}

	if err := h.service.CreateCategory(createdCategory); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (h CategoryHandlers) UpdateCatagory(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedCategory models.Category
	if !utils.BindJSON(c, updatedCategory) {
		return
	}

	if err := h.service.UpdateCatagory(updatedCategory, id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (h CategoryHandlers) DeleteCategory(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Successffuly deleted"})
}
