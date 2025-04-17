package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/models"
	"github.com/tushar27x/portfolioV2-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(ctx *gin.Context, db *gorm.DB) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(user.Passwd), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Passwd = string(hashedPasswd)
	user.Skills = []models.Skill{}
	user.Experience = []models.Experience{}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func LoginUser(ctx *gin.Context, db *gorm.DB) {
	var credentials struct {
		Email  string `json:"email"`
		Passwd string `json:"passwd"`
	}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(credentials.Passwd)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
