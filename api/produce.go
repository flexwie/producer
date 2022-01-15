package api

import (
	"net/http"

	"felixwie.com/producer/logic"
	"felixwie.com/producer/models"
	"github.com/gin-gonic/gin"
)

// Create a new produce
func createProduce(c *gin.Context) {
	var data models.Produce
	if c.BindJSON(&data) != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result, err := logic.Create(data)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get all produces
func getAllProduce(c *gin.Context) {
	result, err := logic.GetAll[models.Produce](nil)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Get one produce
func getOneProduce(c *gin.Context) {
	id := c.Param("id")

	result, err := logic.GetOne[models.Produce](id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
