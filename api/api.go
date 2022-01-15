package api

import (
	"log"
	"net/http"

	"felixwie.com/producer/logic"
	"felixwie.com/producer/models"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()

	produceGroup := router.Group("/produce")
	{
		produceGroup.POST("/", create[models.Produce])
		produceGroup.GET("/", getAll[models.Produce])
		produceGroup.GET("/{id}", getOne[models.Produce])
		produceGroup.DELETE("/{id}", remove[models.Produce])
	}

	inventoryGroup := router.Group("/inventory")
	{
		inventoryGroup.POST("/", create[models.Inventory])
		inventoryGroup.GET("/", getAll[models.Inventory])
		inventoryGroup.GET("/{id}", getOne[models.Inventory])
		inventoryGroup.DELETE("/{id}", remove[models.Inventory])
	}
}

func GetRouter() *gin.Engine {
	return router
}

func getAll[T logic.DbModel](c *gin.Context) {
	result, err := logic.GetAll[T](&logic.QueryOptions{
		Take: 10,
	})
	if err != nil {
		log.Printf("error: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getOne[T logic.DbModel](c *gin.Context) {
	id := c.Param("id")

	result, err := logic.GetOne[T](id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func create[T logic.DbModel](c *gin.Context) {
	var data T
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

func remove[T logic.DbModel](c *gin.Context) {
	id := c.Param("id")

	if logic.Remove[T](id) != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
