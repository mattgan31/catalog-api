package router

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/database"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	db := database.GetDB()
	router := gin.Default()
	ctr := controllers.New(db)

	photoRouter := router.Group("/products")
	{
		photoRouter.POST("/", ctr.CreateProduct)
		photoRouter.PUT("/:productId", ctr.UpdateProduct)
		photoRouter.DELETE("/:productId", ctr.DeleteProduct)
		photoRouter.GET("/", ctr.GetProducts)
		photoRouter.GET("/:productId", ctr.GetProduct)
	}

	return router
}
