package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/models"
	"gorm.io/gorm"
)

func GetExperiencesForUser(ctx *gin.Context, db *gorm.DB) {
	userID, _ := ctx.Get("user_id")

	var experiences []models.Experience
	if err := db.Where("user_id = ?", userID).Find(&experiences).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"experiences": experiences})
}

func AddExperience(ctx *gin.Context, db *gorm.DB) {
	var experiences []models.Experience
	if err := ctx.ShouldBindJSON(&experiences); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("user_id")
	for i := range experiences {
		experiences[i].UserId = userID.(uint)
	}

	if err := db.Create(&experiences).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Experiences added successfully 🎉", "experiences": experiences})
}

func DeleteExperience(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Delete(&models.Experience{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete experience"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})
}

func UpdateExperience(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedExperience models.Experience
	if err := ctx.ShouldBindJSON(&updatedExperience); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingExperience models.Experience
	if err := db.First(&existingExperience, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	existingExperience.CompanyName = updatedExperience.CompanyName
	existingExperience.Designation = updatedExperience.Designation
	existingExperience.Description = updatedExperience.Description

	if err := db.Save(&existingExperience).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Experience updated successfully 🔥", "experience": existingExperience})
}
