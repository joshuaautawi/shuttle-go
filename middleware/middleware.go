package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joshuaautawi/shuttle/auth"
)

func JwtAuthMiddleWare(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorize, err := auth.IsAuthorized(authToken, secret)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"err": err,
				})
				ctx.Abort()
				return
			}
			if authorize {
				userId, err := auth.ExtractIDFromToken(authToken, secret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"err": err,
					})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userId)
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"err": err,
			})
			ctx.Abort()

		}
	}
}
