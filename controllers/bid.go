package controllers

import (
	"auction-system/config"
	"auction-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PlaceBid handles the creation of a new bid
func PlaceBid(c *gin.Context) {
	var bid models.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if bid.AuctionID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction ID is required"})
		return
	}
	if bid.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	if bid.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bid amount must be positive"})
		return
	}

	// Check if auction exists and is active
	var auction models.Auction
	if err := config.DB.First(&auction, bid.AuctionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	// Validate auction time
	now := time.Now()
	if now.Before(auction.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction hasn't started yet"})
		return
	}
	if now.After(auction.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction has already ended"})
		return
	}

	// Check if bid is higher than current highest bid
	var highestBid models.Bid
	config.DB.Where("auction_id = ?", bid.AuctionID).Order("amount desc").First(&highestBid)
	if highestBid.Amount >= bid.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bid amount must be higher than current highest bid"})
		return
	}

	// Set bid time and create the bid
	bid.BidTime = now
	if err := config.DB.Create(&bid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bid)
}

// GetBids retrieves all bids
func GetBids(c *gin.Context) {
	var bids []models.Bid
	if err := config.DB.Find(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}

// GetBidsForAuction retrieves all bids for a specific auction
func GetBidsForAuction(c *gin.Context) {
	auctionID, err := strconv.ParseUint(c.Param("auction_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid auction ID"})
		return
	}

	// Check if auction exists
	var auction models.Auction
	if err := config.DB.First(&auction, auctionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	// Get bids for this auction, ordered by amount (highest first)
	var bids []models.Bid
	if err := config.DB.Where("auction_id = ?", auctionID).Order("amount desc").Find(&bids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bids)
}
