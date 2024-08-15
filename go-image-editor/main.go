package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func getHelloWorld(c *gin.Context) {
	fmt.Println("hello world")
}

func main() {
	r := gin.Default()
	r.GET("/", getHelloWorld)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
