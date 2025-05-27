package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"hnex.com/internal/utils"
)

func AccessTokenMiddleware(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		c.Abort()
	}

	token := strings.Split(authorization, " ")[1]
	claims, err := utils.VerifyToken(token, "JWT_ACCESS_SECRET")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "Unauthorized"})
		c.Abort()
	}

	c.Set("user", claims)
}
