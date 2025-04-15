package controllers

import (
	"auction-system/config"
	"auction-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateWallet creates a new wallet
func CreateWallet(c *gin.Context) {
	var wallet models.Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if wallet.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	if wallet.Balance < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balance cannot be negative"})
		return
	}

	// Set last updated time
	wallet.LastUpdated = time.Now()

	if err := config.DB.Create(&wallet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

// GetWallets retrieves all wallets
func GetWallets(c *gin.Context) {
	var wallets []models.Wallet
	if err := config.DB.Find(&wallets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wallets)
}

// GetWalletByID retrieves a single wallet by ID
func GetWalletByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	var wallet models.Wallet
	if err := config.DB.First(&wallet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// UpdateWallet updates an existing wallet
func UpdateWallet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	var existingWallet models.Wallet
	if err := config.DB.First(&existingWallet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	var updateData struct {
		Balance *float64 `json:"balance"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update balance if provided
	if updateData.Balance != nil {
		if *updateData.Balance < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Balance cannot be negative"})
			return
		}
		existingWallet.Balance = *updateData.Balance
		existingWallet.LastUpdated = time.Now()
	}

	if err := config.DB.Save(&existingWallet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingWallet)
}
