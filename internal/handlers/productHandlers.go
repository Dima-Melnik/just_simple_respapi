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

type ProductHandlers struct {
	service services.ProductServices
}

func NewProductHandlers(s services.ProductServices) *ProductHandlers {
	return &ProductHandlers{service: s}
}

func (h ProductHandlers) GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page > 1 {
		page = 1
	}

	if limit > 10 {
		limit = 10
	}

	offSet := (page - 1) * limit

	products, err := h.service.GetAllProducts(offSet, limit)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h ProductHandlers) GetProductByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h ProductHandlers) CreateProduct(c *gin.Context) {
	var createdProduct models.Product

	if !utils.BindJSON(c, &createdProduct) {
		return
	}

	if err := h.service.CreateProduct(createdProduct); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (h ProductHandlers) UpdateProduct(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedProduct models.Product
	if !utils.BindJSON(c, updatedProduct) {
		return
	}

	if err := h.service.UpdateProduct(updatedProduct, id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h ProductHandlers) DeleteProduct(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		log.Println(err)
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Successfully deleted"})
}
