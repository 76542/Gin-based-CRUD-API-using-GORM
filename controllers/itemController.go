// handles http requests & responses, input validation and routing
// this is where the user request comes first
package controllers

import (
	"example/Crud/logger"
	"example/Crud/models"
	"example/Crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateItem godoc
// @Summary      Create a new item
// @Description  Create a new item with the input payload
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        item body models.Item true "Item to create"
// @Success      201  {object}  models.Item
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /items [post]

func CreateItem(c *gin.Context) {
	var newItem models.Item
	//converting the new item in to a json format and sending it to the server
	if err := c.ShouldBindJSON(&newItem); err != nil {
		logger.CustomLogger.Warn("Invalid request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //if there is an error
		return
	}
	//tries to save the new item in the database using services package
	/*
		services.Item.CreateItem = mockItem.CreateItem
	*/
	createdItem, err := services.Item.CreateItem(&newItem)
	if err != nil {
		logger.CustomLogger.Warn("Error creating item: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdItem)
}

//GetItems godoc
// @Summary      Get all items
// @Description  Get all items
// @Tags         Items
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Item
// @Failure      500  {object}  gin.H
// @Router       /items [get]

func GetItems(c *gin.Context) {
	//fetches the items from the database
	items, err := services.GetItems()
	if err != nil {
		logger.CustomLogger.Warn("Error fetching items: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

//GetItembyID godoc
// @Summary      Get item by ID
// @Description  Get an item by its ID
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 id path int true "Item ID"
// @Success      200  {array}  models.Item
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Router       /items/{id} [get]

func GetItemByID(c *gin.Context) {
	//extracts the id from the url
	idParam := c.Param("id")
	//converting the id(string) to an unsigned integer
	id, err := strconv.ParseUint(idParam, 10, 32)
	//agar item id wrong hai
	if err != nil {
		logger.CustomLogger.Warn("Invalid item id: ", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}
	//agar item hi nahi mila
	item, err := services.GetItemByID(uint(id))
	if err != nil {
		logger.CustomLogger.Warn("Error fetching item by id: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	//agar item mila
	c.JSON(http.StatusOK, item)
}

//UpdateItem godoc
// @Summary      Get item by ID
// @Description  Get an item by its ID
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 id path int true "Item ID"
// @Param        item body models.Item true "Item to update"
// @Success      200  {object}  models.Item
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Router       /items/{id} [put]

func UpdateItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	//agar id is wrong
	if err != nil {
		logger.CustomLogger.Warn("Invalid item id: ", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}
	//this new variable holds the updated info
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		logger.CustomLogger.Warn("Invalid request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//item update karne ke liye
	item, err := services.UpdateItem(uint(id), &updatedItem)
	//agar item nahi mila
	if err != nil {
		logger.CustomLogger.Error("Error updating item: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

//DeleteItem godoc
// @Summary			Delete an item
// @Description		Delete an item by its ID
// @Tags			Items
// @Accept			json
// @Produce			json
// @Param			id path int true "Item ID"
// @Success			204
// @Failure			400  {object}  gin.H
// @Failure			404  {object}  gin.H
// @Router			/items/{id} [delete]

func DeleteItem(c *gin.Context) {
	//deletes item by id so to get id,
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32) //converts
	//agar id wrong hai
	if err != nil {
		logger.CustomLogger.Warn("Invalid item id: ", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}

	if err := services.DeleteItem(uint(id)); err != nil {
		logger.CustomLogger.Error("Error deleting item: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) //agar item nahi mila
		return
	}
	c.JSON(http.StatusNoContent, nil) //deletion was successful and returns no content

}
