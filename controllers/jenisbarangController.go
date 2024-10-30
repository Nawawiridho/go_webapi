package controllers

import (
	"fmt"
	"gin/database"
	"gin/models"
	"gin/respon"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Mendapatkan semua barang
func GetJenisBarang(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id_jenis, nama_jenis FROM jenisbarang")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var jnsbrg_ []models.Jenisbarang
	for rows.Next() {
		var jbrg models.Jenisbarang
		if err := rows.Scan(&jbrg.ID_jenis, &jbrg.Nama_jenis); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		jnsbrg_ = append(jnsbrg_, jbrg)
	}

	c.JSON(http.StatusOK, jnsbrg_)
}

// Menambahkan barang baru
func CreateJenisBarang(c *gin.Context) {
	var jbrg models.Jenisbarang

	if err := c.ShouldBindJSON(&jbrg); err != nil {
		fmt.Println("Error tipe data JSON:", err) // Debugging output
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO jenisbarang (id_jenis, nama_jenis) VALUES (?, ?)", jbrg.ID_jenis, jbrg.Nama_jenis)
	if err != nil {
		fmt.Println("Error saat insert:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusCreated, respon.Response{
		Status:  respon.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"barang created"},
	})
}

// Mengupdate barang
func UpdateJenisBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error converting ID:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	var jbrg models.Jenisbarang
	if err := c.ShouldBindJSON(&jbrg); err != nil {
		fmt.Println("Error tipe data JSON:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("UPDATE jenisbarang SET nama_jenis = ? WHERE id_jenis = ?", jbrg.Nama_jenis, id)
	if err != nil {
		fmt.Println("Error saat update:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusCreated, respon.Response{
		Status:  respon.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"barang updated"},
	})
}

// Menghapus barang
func DeleteJenisBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("DELETE FROM jenisbarang WHERE id_jenis = ?", id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusCreated, respon.Response{
		Status:  respon.StatusInfo{Code: http.StatusCreated, Error: 0, Message: "OK"},
		Results: []string{"barang deleted"},
	})
}
