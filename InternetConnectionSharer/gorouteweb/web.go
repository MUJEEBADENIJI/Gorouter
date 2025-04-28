package gorouteweb

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Load all templates (including components like sidebar and footer)
	router.LoadHTMLGlob("gorouteweb/templates/*.html")

	// Serve static files (CSS, JS, images)
	router.Static("/static", "./gorouteweb/static")

	// Common data (example: CurrentYear for footer)
	commonData := gin.H{
		"CurrentYear": time.Now().Year(),
	}

    
	// Routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", commonData)
	})

	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", commonData)
	})

	router.GET("/overview", func(c *gin.Context) {
		c.HTML(200, "overview.html", commonData)
	})

	router.GET("/Devicemanagement", func(c *gin.Context) {
		c.HTML(200, "Devicemanagement.html", commonData)
	})
}
