package main

import (
    "github.com/gin-gonic/gin"
    "example.com/InternetConnectionSharer/gorouteweb" // Import your web package
)

func main() {
    r := gin.Default()

    // Set trusted proxies to only allow localhost
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // Register routes from the gorouteweb package
    gorouteweb.RegisterRoutes(r)

    // Start the server
    r.Run(":5005") // Listen on port 5005
}
