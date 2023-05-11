package main

import (
	"github.com/gin-gonic/gin"
)

func getTemplate(location string, weather *weatherData) gin.H {
	return gin.H{
		"location": location,
		"weather":  weather,
	}
}
