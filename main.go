package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	tasks := g.Group("/task")
	tasks.GET("/:id")
	tasks.POST("/")
	tasks.PUT("/:id")

	g.Run()
}
