package gorouteweb

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    // Load both templates and components
    router.LoadHTMLGlob("gorouteweb/templates/*.html")
    
    // Serve static files (CSS, JS, images)
    router.Static("/static", "./gorouteweb/static")


    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })


    router.GET("/dashboard", func(c *gin.Context) {
        c.HTML(200, "dashboard.html", gin.H{})
    })


    router.GET("/overview", func(c *gin.Context) {
        c.HTML(200, "overview.html", gin.H{})
    })


    router.GET("/Devicemanagement", func(c *gin.Context) {
        c.HTML(200, "Devicemanagement.html", gin.H{})
    })
}
