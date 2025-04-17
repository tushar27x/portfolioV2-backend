package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/models"
	"gorm.io/gorm"
)

func GetProjectsForUser(ctx *gin.Context, db *gorm.DB) {
	userID, _ := ctx.Get("user_id")

	var projects []models.Project
	if err := db.Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

func AddProject(ctx *gin.Context, db *gorm.DB) {
	var projects []models.Project

	if err := ctx.ShouldBindJSON(&projects); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("user_id")
	for i := range projects {
		projects[i].UserId = userID.(uint)
	}
	if err := db.Create(&projects).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project added successfully",
		"project": projects,
	})
}

func UpdateProject(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedProject, existingProject models.Project
	if err := ctx.ShouldBindJSON(&updatedProject); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&existingProject, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "project does not exist"})
		return
	}

	existingProject.Title = updatedProject.Title
	existingProject.Description = updatedProject.Description
	existingProject.GithubLink = updatedProject.GithubLink
	existingProject.LiveLink = updatedProject.LiveLink
	existingProject.Stack = updatedProject.Stack

	if err := db.Save(&existingProject).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
		"project": existingProject,
	})
}

func DeleteProject(ctx *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Delete(&models.Project{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project Deleted successfully",
	})
}
