package main

import (

	"github.com/gin-gonic/gin"
	"github.com/yhlin66/go-side/controller"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/", controller.HomepageHandler)
	server.GET("/index")

	server.Run()
}
