package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/upmaru-stage/gemini/initializers"
	"github.com/upmaru-stage/gemini/models"
)

type CreateParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePost(c *gin.Context) {
	var createParams CreateParams
	c.BindJSON(&createParams)

	post := models.Post{Title: createParams.Title, Body: createParams.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func ListPosts(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ShowPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
