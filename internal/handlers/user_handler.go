package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hnex.com/internal/services"
	"hnex.com/internal/utils"
)

type UserHandler struct {
	Service services.UserService
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	_id, err := utils.ConvertID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	user, err := h.Service.Repository.FindById(_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": user,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		return
	}

	claims, ok := user.(utils.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		return
	}

	user, err := h.Service.Repository.FindById(claims.Sub)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": user,
	})
}
