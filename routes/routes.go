package routes

import (
	"github.com/dario-nihil/go_rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register")

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
