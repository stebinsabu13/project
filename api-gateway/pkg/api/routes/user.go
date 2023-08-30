package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api/handler"
)

func RegisterRoutes(api *gin.RouterGroup, userHandler handler.UserHandler, methodHandler handler.MethodHandler) {
	user := api.Group("/user")
	{
		user.POST("/", userHandler.CreateUser)
		user.GET("/", userHandler.GetUserById)
		user.PATCH("/", userHandler.UpdateUser)
		user.DELETE("/", userHandler.DeleteUser)
	}
	method := api.Group("/methods")
	{
		method.POST("/", methodHandler.Methods)
	}
}
