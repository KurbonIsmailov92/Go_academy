package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-crud/logger"
	"todo-crud/pkg/service"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
)

func checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		logger.Info.Printf("[controller-midlwere] empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		logger.Info.Printf("[controller-midlwere] invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		logger.Info.Printf("[controller-midlwere] token is empty")
		return
	}

	accessToken := headerParts[1]

	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.Info.Printf("[controller-midlwere] invalid token")
		return
	}
	c.Set(userIDCtx, claims.UserID)
	c.Next()
}
