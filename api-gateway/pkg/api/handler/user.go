package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	client "github.com/stebin13/x-tentioncrew/api-gateway/pkg/client/interfaces"
	"github.com/stebin13/x-tentioncrew/api-gateway/pkg/utils"
)

type UserHandler struct {
	UserClient client.UserClient
}

func NewUserHandler(client client.UserClient) UserHandler {
	return UserHandler{
		UserClient: client,
	}
}

func (cr *UserHandler) CreateUser(c *gin.Context) {
	var body utils.UserDetails
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error" + err.Error(),
		})
		return
	}
	log.Println("inside handler", body)
	res, err := cr.UserClient.UserCreate(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"responseid": &res.Responseid,
	})
}

func (cr *UserHandler) GetUserById(c *gin.Context) {
	userid := c.Query("id")
	res, err := cr.UserClient.GetUserById(context.Background(), userid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"user": res.User,
	})
}

func (cr *UserHandler) UpdateUser(c *gin.Context) {
	userid, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid id" + err1.Error(),
		})
		return
	}
	var body utils.UserDetails
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "binding error" + err.Error(),
		})
		return
	}
	res, err := cr.UserClient.UpdateUser(context.Background(), userid, body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"user": res.User,
	})
}

func (cr *UserHandler) DeleteUser(c *gin.Context) {
	userid := c.Query("id")
	res, err := cr.UserClient.DeleteUser(context.Background(), userid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"Success": "User deleted",
	})
}
