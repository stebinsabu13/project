package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api/handler"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/api/routes"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/config"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(c *config.Config, userHandler handler.UserHandler, methodHandler handler.MethodHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	routes.RegisterRoutes(engine.Group("/"), userHandler, methodHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})
	log.Println("inside server", c.Port)
	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.Run(c.Port)
}
