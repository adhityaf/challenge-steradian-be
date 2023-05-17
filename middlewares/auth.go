package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/adhityaf/challenge-steradian-be/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Token Not Found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Bearer Not Found",
			})
			return
		}

		tokenStr := strings.Split(token, "Bearer ")[1]
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Token Not Found",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)
		userId := data["id"].(float64)
		strUserId := strconv.FormatFloat(userId, 'f', -1, 64)

		ctx.Set("id", strUserId)
		ctx.Set("name", data["name"])
		ctx.Set("role", data["role"])
		ctx.Set("expire", data["expire"])
		ctx.Set("expire_date", data["expire_date"])

		if data["expire_date"] == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
			return
		}

		timeNow := time.Now()
		expiredTime := data["expire_date"].(string)

		parsed, _ := time.Parse(time.RFC3339, expiredTime)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		if timeNow.After(parsed) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Token has expired, please login again",
			})
			return
		}

		ctx.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Your role is not allowed to use this API",
			})
			return
		}
	}
}
