package middlewares

import (
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("Authorization")
		if token == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)

		if err !=nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}