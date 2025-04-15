package controllers

import (
	"auction-system/config"
	"auction-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAuction creates a new auction
func CreateAuction(c *gin.Context) {
	var auction models.Auction
	if err := c.ShouldBindJSON(&auction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate auction times
	if auction.EndTime.Before(auction.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
		return
	}

	if auction.StartTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start time must be in the future"})
		return
	}

	if err := config.DB.Create(&auction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, auction)
}

// GetAuctions retrieves all auctions
func GetAuctions(c *gin.Context) {
	var auctions []models.Auction
	if err := config.DB.Find(&auctions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, auctions)
}

// GetAuctionByID retrieves a single auction by ID
func GetAuctionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid auction ID"})
		return
	}

	var auction models.Auction
	if err := config.DB.First(&auction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	c.JSON(http.StatusOK, auction)
}

// UpdateAuction updates an existing auction
func UpdateAuction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid auction ID"})
		return
	}

	var existingAuction models.Auction
	if err := config.DB.First(&existingAuction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	var updateData models.Auction
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate auction times if they're being updated
	if !updateData.StartTime.IsZero() || !updateData.EndTime.IsZero() {
		startTime := existingAuction.StartTime
		endTime := existingAuction.EndTime

		if !updateData.StartTime.IsZero() {
			startTime = updateData.StartTime
		}
		if !updateData.EndTime.IsZero() {
			endTime = updateData.EndTime
		}

		if endTime.Before(startTime) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
			return
		}
	}

	// Update only non-zero fields
	if updateData.ItemID != 0 {
		existingAuction.ItemID = updateData.ItemID
	}
	if updateData.StartingPrice != 0 {
		existingAuction.StartingPrice = updateData.StartingPrice
	}
	if !updateData.StartTime.IsZero() {
		existingAuction.StartTime = updateData.StartTime
	}
	if !updateData.EndTime.IsZero() {
		existingAuction.EndTime = updateData.EndTime
	}

	if err := config.DB.Save(&existingAuction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAuction)
}

// DeleteAuction removes an auction from the database
func DeleteAuction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid auction ID"})
		return
	}

	var auction models.Auction
	if err := config.DB.First(&auction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	if err := config.DB.Delete(&auction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auction deleted successfully"})
}
