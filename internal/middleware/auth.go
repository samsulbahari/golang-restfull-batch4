package middleware

import (
	"context"
	"restfull/internal/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

var bearer = "Bearer "

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}

		//untuk cek apakah header = bearer
		if !strings.HasPrefix(authHeader, bearer) {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}
		token := strings.Split(authHeader, " ")
		user := domain.User{}
		data, err := user.DecryptJwt(token[1])
		if err != nil {
			ctx.JSON(401, gin.H{
				"code":    "401",
				"message": "Invalid token",
			})

			ctx.Abort()
			return
		}
		userID := int(data["user_id"].(float64))
		ctxUserID := context.WithValue(ctx.Request.Context(), "user_id", userID)
		ctx.Request = ctx.Request.WithContext(ctxUserID)
		ctx.Next()
	}
}
