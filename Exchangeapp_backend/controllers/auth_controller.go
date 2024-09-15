package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context){
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err !=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err:= utils.HashPassword(user.Password)

	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)

	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err:= global.Db.AutoMigrate(&user); err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err:= global.Db.Create(&user).Error; err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context){
	var input struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err:= global.Db.Where("username = ?", input.Username).First(&user).Error; err !=nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}	

	if !utils.CheckPassword(input.Password, user.Password){
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)

	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}