package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/upmaru-stage/gemini/controllers"
	"github.com/upmaru-stage/gemini/initializers"
	"github.com/upmaru-stage/gemini/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func web() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.ListPosts)
	r.GET("/posts/:id", controllers.ShowPost)

	r.Run()
}

func migrate() {
	fmt.Println("INFO: Running migrations")

	initializers.DB.AutoMigrate(&models.Post{})
}

func main() {
	mode := os.Args[1]

	switch mode {
	case "web":
		web()
	case "migrate":
		migrate()
	default:
		web()
	}
}
