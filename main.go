package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	apikey := os.Getenv("OWM_API_KEY")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.POST("/", func(c *gin.Context) {
		location := c.PostForm("location")
		if location == "" {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"message": "error",
			})
			return
		}
		output, err := getWeather(apikey, location)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"message": "failed to get weather data",
			})
			return
		}
		c.HTML(http.StatusOK, "interface.html", getTemplate(location, output))
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "interface.html", gin.H{})
	})

	router.Run(":8080")
}
