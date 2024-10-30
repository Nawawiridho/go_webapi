package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/items", controllers.GetItems)
	router.POST("/items", controllers.CreateItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)
}

func BarangRoutes(router *gin.Engine) {
	router.GET("/getall", controllers.GetAll)
	router.GET("/get_barang", controllers.GetBarang)
	router.POST("/simpan_barang", controllers.CreateBarang)
	router.PUT("/update_barang/:id", controllers.UpdateBarang)
	router.DELETE("/hapus_barang/:id", controllers.DeleteBarang)
}

func JenisBarangRoutes(router *gin.Engine) {
	router.GET("/get_jenisbarang", controllers.GetJenisBarang)
	router.POST("/simpan_jenisbarang", controllers.CreateJenisBarang)
	router.PUT("/update_jenisbarang/:id", controllers.UpdateJenisBarang)
	router.DELETE("/hapus_jenisbarang/:id", controllers.DeleteJenisBarang)
}
