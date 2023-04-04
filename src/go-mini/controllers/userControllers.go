package controllers

import (
	"net/http"

	"github.com/Emc002/go-mini/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Index(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) FilterByJob(c *gin.Context) {
	job := c.Param("job")

	var users []models.User

	if err := models.DB.Where("job = ?", job).Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}
