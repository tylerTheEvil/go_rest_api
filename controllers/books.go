package controllers

import (
	"net/http"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	context.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(context *gin.Context) {
	var input CreateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	context.JSON(http.StatusCreated, gin.H{"data": book})
}

func FindBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&book)
	context.JSON(http.StatusOK, gin.H{"data": true})
}
