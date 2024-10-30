package controllers

import (
	"gin/database"
	"gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Mendapatkan semua item

func GetItems(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, price FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

// Menambahkan item baru
func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO items (name, price) VALUES (?, ?)", item.Name, item.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Status:  models.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"Item created"},
	})
}

// Mengupdate item
func UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("UPDATE items SET name = ?, price = ? WHERE id = ?", item.Name, item.Price, id)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Item updated"})
	c.JSON(http.StatusCreated, models.Response{
		Status:  models.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"Item updated"},
	})
}

// Menghapus item
func DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	c.JSON(http.StatusCreated, models.Response{
		Status:  models.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"Item deleted"},
	})
}
