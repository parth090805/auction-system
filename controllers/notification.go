package controllers

import (
    "auction-system/config"
    "auction-system/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SendNotification(c *gin.Context) {
    var notif models.Notification
    if err := c.ShouldBindJSON(&notif); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&notif)
    c.JSON(http.StatusOK, notif)
}

func GetNotifications(c *gin.Context) {
    var notifs []models.Notification
    config.DB.Find(&notifs)
    c.JSON(http.StatusOK, notifs)
}
