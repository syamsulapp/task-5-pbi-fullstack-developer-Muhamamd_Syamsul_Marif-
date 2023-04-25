package main

import "github.com/gin-gonic/gin"

func dummyJson(res *gin.Context) {
	var arr []int
	res.JSON(200, gin.H{
		"message": append(arr, 1, 2, 3),
	})
}

func main() {
	req := gin.Default()
	req.POST("/dummy", dummyJson)
	req.Run()
}
