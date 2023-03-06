package productController

import (
	"net/http"
	"strconv"

	"github.com/Emc002/goGin-RestAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id= ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data cannot Update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data has been Updated", "data": product})
}

func Delete(c *gin.Context) {
	var product models.Product
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data cannot Delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messsage": "Data has been Deleted"})
}
