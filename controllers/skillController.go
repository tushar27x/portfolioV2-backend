package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/models"
	"gorm.io/gorm"
)

func GetSkillsForUser(ctx *gin.Context, db *gorm.DB) {
	userID, _ := ctx.Get("user_id")

	var skills []models.Skill
	if err := db.Where("user_id = ?", userID).Find(&skills).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"skills": skills})
}

func AddSkill(ctx *gin.Context, db *gorm.DB) {
	var skills []models.Skill

	if err := ctx.ShouldBindJSON(&skills); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("user_id")
	for i := range skills {
		skills[i].UserId = userID.(uint)
	}

	for _, skill := range skills {
		var existing models.Skill
		if err := db.Where("name = ? AND user_id = ?", skill.Name, skill.UserId).First(&existing).Error; err == nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Skill '" + skill.Name + "' already exists for this user"})
			return
		}
	}

	if err := db.Create(&skills).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Skill added successfully 🎉🎉", "skills": skills})
}

func DeleteSkill(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	result := db.Unscoped().Delete(&models.Skill{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete skill"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}

func UpdateSkill(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedSkill models.Skill
	if err := ctx.ShouldBindJSON(&updatedSkill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingSkill models.Skill
	if err := db.First(&existingSkill, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedSkill.Score < 1 || updatedSkill.Score > 10 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 1 and 10"})
		return
	}

	existingSkill.Name = updatedSkill.Name
	existingSkill.Score = updatedSkill.Score
	existingSkill.ImageURL = updatedSkill.ImageURL

	if err := db.Save(&existingSkill).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Skill updated successfully 🔥",
		"skill":   existingSkill,
	})
}
