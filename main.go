package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// Handle Index
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}
