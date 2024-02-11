package middleware

import (
	"example.com/ideanest-task/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader,"Bearer ") || len(strings.Split(authHeader, " ")) != 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	token := strings.Split(authHeader, " ")[1]
	claims, err := utils.ExtractClaims(token)

	if err != nil || claims.Type != "access"{

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		ctx.Abort()
		return
	}
	ctx.Set("email",claims.Email)
	ctx.Next()
}
