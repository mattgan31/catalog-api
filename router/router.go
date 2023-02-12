package router

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	photoRouter := router.Group("/products")
	{
		photoRouter.POST("/", controllers.CreateProduct)
		photoRouter.PUT("/:productId", controllers.UpdateProduct)
		photoRouter.DELETE("/:productId", controllers.DeleteProduct)
		photoRouter.GET("/", controllers.GetProducts)
		photoRouter.GET("/:productId", controllers.GetProduct)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controllers.UserRegister)

	}

	return router
}
