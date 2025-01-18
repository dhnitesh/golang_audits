package handlers

import (
	"net/http"

	"logger/internal/models"
	"logger/internal/mongo"

	"github.com/gin-gonic/gin"
)

// DataHandler handles incoming data requests
func DataHandler(c *gin.Context) {
	var data models.Data
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	if err := mongo.SaveToMongo(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data saved successfully"})
}

// GetDataHandler returns all data from MongoDB that are paginated which a numer that is passed as a query parameter
func GetDataHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	data, err := mongo.GetDataFromMongo(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get data"})
		return
	}

	c.JSON(http.StatusOK, data)
}
