package router

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	productRouter := router.Group("/products")
	{
		// productRouter.Use(middleware.Authentication())
		productRouter.POST("/", middleware.Authentication(), controllers.CreateProduct)
		productRouter.PUT("/:productId", middleware.Authentication(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middleware.Authentication(), controllers.DeleteProduct)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:productId", controllers.GetProduct)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

	}

	return router
}
