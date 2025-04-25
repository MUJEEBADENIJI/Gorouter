package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})


//everytime you want to add a new page
//u  must register the route like this

r.GET("/dashboard", func(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{})
})

r.GET("/overview", func(c *gin.Context) {
	c.HTML(200, "overview.html", gin.H{})
})

r.GET("/Devicemanagement", func(c *gin.Context) {
	c.HTML(200, "Devicemanagement.html", gin.H{})
})










//do not touch this..... be careful please .
	r.Run(":5005")
}
