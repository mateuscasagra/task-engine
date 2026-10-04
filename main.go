package main

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, "Hello World!")
}

func main() {
	g := gin.Default()
	g.GET("/hello", HelloWorld)
	g.Run()

}
