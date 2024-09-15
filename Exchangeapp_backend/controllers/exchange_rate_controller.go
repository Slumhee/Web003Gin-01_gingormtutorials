package controllers

import (
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRate(ctx *gin.Context){
	var exchangeRate models.ExchangeRate

	if err:= ctx.ShouldBindJSON(&exchangeRate); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exchangeRate.Date = time.Now()

	if err := global.Db.AutoMigrate(&exchangeRate); err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&exchangeRate).Error; err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, exchangeRate)
}

func GetExchangeRates(ctx *gin.Context){
	var exchangeRates []models.ExchangeRate

	if err:= global.Db.Find(&exchangeRates).Error; err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}