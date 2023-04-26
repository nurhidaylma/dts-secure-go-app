package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-secure-go-app/controllers"
	"github.com/nurhidaylma/dts-secure-go-app/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	userRouter := r.Group("/api/v1/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
