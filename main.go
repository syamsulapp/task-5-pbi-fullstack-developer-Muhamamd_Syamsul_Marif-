package main

import (
	"github.com/gin-gonic/gin"
)

// method for create response api mahasiswa
func mahasiswa(res *gin.Context) {
	res.JSON(200, gin.H{
		"nama": "samsul",
		"nim":  "f1b118038",
	})
}

func universitas(res *gin.Context) {
	res.JSON(200, gin.H{
		"name":   "universitas haluoleo",
		"alamat": "jalan hea mokodompit",
	})
}

func fakultas(res *gin.Context) {
	arr := [5]int{1, 2, 3, 4, 5}

	res.JSON(200, gin.H{
		"message": "Successfully Data",
		"data":    arr,
	})
}

func main() {
	req := gin.Default()
	req.GET("/fakultas", fakultas)
	req.GET("/universitas", universitas)
	req.GET("/mahasiswa", mahasiswa)
	req.Run()

}
