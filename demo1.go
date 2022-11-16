package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	n := gin.Default()
	//for i := 1; i <= 10; i++ {
	n.GET("/table", func(c *gin.Context) {
		for i := 1; i <= 10; i++ {
			////fmt.Println("2 * ", i, " = ", 2*i)
			//c.JSON(http.StatusOK, gin.H{
			//fmt.Println("2 * ", i, " = ", 2*i),
			//"table of 2 on terminal": "",

		}
		c.JSON(http.StatusOK, gin.H{
			"table of 2 ": "on terminal",
			"fmt.Println("2 * ",i ,"=", "2*i")",
		})

	})
	n.Run()
}
