package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type MethodHandler struct {
	MethodClient interfaces.MethodCli
}

func NewMethodHandler(client interfaces.MethodCli) MethodHandler {
	return MethodHandler{
		MethodClient: client,
	}
}

func (cr *MethodHandler) Methods(c *gin.Context) {
	var body utils.MethodReq
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error" + err.Error(),
		})
		return
	}
	res, err := cr.MethodClient.Methods(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"users": res.User,
	})
}
