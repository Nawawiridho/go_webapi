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
func GetAll(c *gin.Context) {
	rows, err := database.DB.Query("SELECT a.id_transaksi, b.nama_barang, c.nama_jenis AS jenis_barang, a.jumlah_terjual, a.tanggal_transaksi FROM transaksi a JOIN barang b ON a.id_barang = b.id_barang JOIN jenisbarang c ON b.id_jenis = c.id_jenis")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.GetAll
	for rows.Next() {
		var item models.GetAll
		if err := rows.Scan(&item.NO, &item.Nama_barang, &item.Jenis_barang, &item.Jumlah_terjual, &item.Tanggal_transaksi); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func GetBarang(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id_barang, nama_barang, stok, id_jenis FROM barang")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var brg_ []models.Barang
	for rows.Next() {
		var brg models.Barang
		if err := rows.Scan(&brg.ID_barang, &brg.Nama_barang, &brg.Stok, &brg.ID_jenis); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		brg_ = append(brg_, brg)
	}

	c.JSON(http.StatusOK, brg_)
}

// Menambahkan barang baru
func CreateBarang(c *gin.Context) {
	var brg models.Barang

	if err := c.ShouldBindJSON(&brg); err != nil {
		fmt.Println("Error tipe data JSON:", err) // Debugging output
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO barang (nama_barang, stok, id_jenis) VALUES (?, ?, ?)", brg.Nama_barang, brg.Stok, brg.ID_jenis)
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
func UpdateBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error converting ID:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  models.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	var brg models.Barang
	if err := c.ShouldBindJSON(&brg); err != nil {
		fmt.Println("Error tipe data JSON:", err) // Debugging output
		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("UPDATE barang SET nama_barang = ?, stok = ? WHERE id_barang = ?", brg.Nama_barang, brg.Stok, id)
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
func DeleteBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		c.JSON(http.StatusInternalServerError, respon.Response{
			Status:  respon.StatusInfo{Code: http.StatusInternalServerError, Error: 1, Message: err.Error()},
			Results: nil,
		})
		return
	}

	_, err = database.DB.Exec("DELETE FROM barang WHERE id_barang = ?", id)
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
