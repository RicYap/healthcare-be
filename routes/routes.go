package routes

import (
	"healthcare-be/controllers"
	"healthcare-be/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/auth/signup", controllers.SignUp)
	r.POST("/auth/signin", controllers.SignIn)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/lab-results", controllers.GetLabResults)
	auth.POST("/lab-results", controllers.AddLabResult)
}
