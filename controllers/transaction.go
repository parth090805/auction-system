package controllers

import (
    "auction-system/config"
    "auction-system/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateTransaction(c *gin.Context) {
    var tx models.Transaction
    if err := c.ShouldBindJSON(&tx); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&tx)
    c.JSON(http.StatusOK, tx)
}

func GetTransactions(c *gin.Context) {
    var txs []models.Transaction
    config.DB.Find(&txs)
    c.JSON(http.StatusOK, txs)
}
