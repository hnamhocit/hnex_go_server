package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Db *gorm.DB
}

func (h *AuthHandler) Register(c *gin.Context) {}

func (h *AuthHandler) Login(c *gin.Context) {}

func (h *AuthHandler) Logout(c *gin.Context) {}

func (h *AuthHandler) RefreshToken(c *gin.Context) {}
