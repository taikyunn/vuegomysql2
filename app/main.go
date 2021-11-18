package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/new", posting)
	r.Run(":8000")
}

func posting(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
}
