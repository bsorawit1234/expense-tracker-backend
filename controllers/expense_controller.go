package controllers

import (
	"net/http"
	"strconv"

	"github.com/bsorawit1234/expense-tracker-backend/models"
	"github.com/gin-gonic/gin"
)

func GetExpenses(c *gin.Context) {
	userID, _ := c.Get("userID")
	var expenses []models.Expense

	if err := models.DB.Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func CreateExpense(c *gin.Context) {
	var expense models.Expense
	userID, _ := c.Get("userID")
	expense.UserID = userID.(uint)

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func UpdateExpense(c *gin.Context) {
	var expense models.Expense
	userID, _ := c.Get("userID")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ? AND user_id ?", id, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.save(&expense)
	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	var expense models.Expense
	userID, _ := c.Get("userID")
	id, _ := strconv.Atoi(c.Param("id"))

	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	models.DB.Delete(&expense)
	c.JSON(http.StatusOK, gin.H{"message": "Expenese deleted"})
}
