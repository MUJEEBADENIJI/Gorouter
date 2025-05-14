package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "example.com/InternetConnectionSharer/nat"

)

// StartHotspotHandler is called when the Start button is clicked.
func StartHotspotHandler(c *gin.Context) {
    err := nat.StartHotspot()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Hotspot started!",
    })
}

// StopHotspotHandler (optional)
func StopHotspotHandler(c *gin.Context) {
    err := nat.StopHotspot()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Hotspot stopped!",
    })
}
