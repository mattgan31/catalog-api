package controllers

import (
	"net/http"
	"strconv"
	"time"

	"example.com/m/v2/database"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) CreateProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Product.Created_At = time.Now()
	Product.Updated_At = time.Now()

	err := db.Debug().Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":           Product.ID,
		"product_name": Product.Product_Name,
		"description":  Product.Description,
		"image":        Product.Image,
		"created_at":   Product.Created_At,
	})

}

func (idb *InDB) UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))

	newProduct := models.Product{
		Product_Name: Product.Product_Name,
		Description:  Product.Description,
		Image:        Product.Image,
	}

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Product.ID = uint(productId)

	Product.Updated_At = time.Now()
	err := db.First(&Product, productId).Updates(&newProduct).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           Product.ID,
		"product_name": Product.Product_Name,
		"description":  Product.Description,
		"image":        Product.Image,
		"updated_at":   Product.Updated_At,
	})
}

func (idb *InDB) DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Model(&Product).Where("id = ?", productId).Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Successfully Deleted",
	})
}

func (idb *InDB) GetProduct(c *gin.Context) {
	db := database.GetDB()
	var Product models.Product

	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.First(&Product, productId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": Product,
	})
}

func (idb *InDB) GetProducts(c *gin.Context) {
	db := database.GetDB()
	Product := []models.Product{}
	var result gin.H

	err := db.Find(&Product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	if len(Product) <= 0 {
		result = gin.H{
			"result": "data not found",
			"code":   200,
		}
	} else {
		result = gin.H{
			"result": Product,
			"code":   200,
		}
	}

	c.JSON(http.StatusOK, result)
}
