package main

import (
	"CompareGo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // load all the html files in the templates folder
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is working",
		})
	})
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		amazon_result := services.GetAmazon(query)
		flipkart_result := services.GetFlipkart(query)
		c.JSON(200, gin.H{
			"query":    query,
			"amazon":   amazon_result,
			"flipkart": flipkart_result,
		})
	})

	r.Run() // listen and serve on http://localhost:8080/
}
