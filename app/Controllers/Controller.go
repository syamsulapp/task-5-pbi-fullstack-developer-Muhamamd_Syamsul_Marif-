package main

import "github.com/gin-gonic/gin"

func dummyJson(res *gin.Context) {
	var arr = []string{"apple", "banana", "mango", "durian", "pineapple"}
	res.JSON(200, gin.H{
		"data": append(arr[:3], "rambutan"),
	})
}

func main() {
	req := gin.Default()
	req.POST("/dummy", dummyJson)
	req.Run()
}
