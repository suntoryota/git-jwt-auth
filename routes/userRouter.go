package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/suntoryota/go-jwt-auth/controllers"
	"github.com/suntoryota/go-jwt-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
