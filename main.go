package main

import (
	"github.com/gin-gonic/gin"
	"github.com/task-engine/src"
	"github.com/task-engine/src/Controller"
)

func main() {
	g := gin.Default()
	db := src.Conn()
	task := Controller.TaskController{
		Session: db,
	}
	defer db.Close()
	tasks := g.Group("/task", task.GetTask)
	tasks.GET("/:id")
	tasks.POST("/", task.CreateTask)
	tasks.PUT("/:id")

	g.Run()
}
