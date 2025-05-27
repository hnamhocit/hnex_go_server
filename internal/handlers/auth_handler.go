package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	dtos "hnex.com/internal/dtos/auth"
	"hnex.com/internal/models"
	"hnex.com/internal/services"
	"hnex.com/internal/utils"
)

type AuthHandler struct {
	Service services.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user dtos.RegisterDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	existingUser, err := h.Service.UserRepository.FindByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "User already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	userModel := &models.User{
		Email:       user.Email,
		Password:    hashedPassword,
		DisplayName: user.DisplayName,
	}

	err = h.Service.Repository.CreateUser(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(int32(userModel.ID), userModel.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	err = h.Service.UpdateRefreshToken(int32(userModel.ID), &refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "Success", "data": gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user dtos.LoginDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	existingUser, err := h.Service.UserRepository.FindByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "Email not found"})
		return
	}

	match, err := utils.VerifyPassword(user.Password, existingUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "Password is incorrect"})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "Password is incorrect"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(int32(existingUser.ID), existingUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	err = h.Service.UpdateRefreshToken(int32(existingUser.ID), &refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "Success", "data": gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		return
	}

	userID := claims.(*utils.JWTClaims).Sub

	err := h.Service.UpdateRefreshToken(userID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "Success"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var refreshToken dtos.RefreshTokenDto
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		return
	}

	userID := claims.(*utils.JWTClaims).Sub
	existingUser, err := h.Service.UserRepository.FindById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "User not found"})
		return
	}

	match, err := utils.VerifyPassword(refreshToken.RefreshToken, existingUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "Refresh token is incorrect"})
		return
	}

	accessToken, newRefreshToken, err := utils.GenerateTokens(int32(existingUser.ID), existingUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	err = h.Service.UpdateRefreshToken(int32(existingUser.ID), &newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "Success", "data": gin.H{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	}})
}
