package gorouteweb

import (
	"time"
	"example.com/InternetConnectionSharer/gorouteweb/handlers"
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

	router.GET("/devicemanagement", func(c *gin.Context) {
		c.HTML(200, "devicemanagement.html", commonData)
	})

	router.GET("/security", func(c *gin.Context) {
		c.HTML(200, "security.html", commonData)
	})

	router.GET("/log", func(c *gin.Context) {
		c.HTML(200, "log.html", commonData)
	})

	// ✅ New API route for starting hotspot
	 router.POST("/api/start-hotspot", handlers.StartHotspotHandler)
	 // Optional stop:
	 router.POST("/api/stop-hotspot", handlers.StopHotspotHandler)
}
