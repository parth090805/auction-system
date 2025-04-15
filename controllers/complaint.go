package controllers

import (
	"auction-system/config"
	"auction-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FileComplaint handles the creation of a new complaint
func CreateComplaint(c *gin.Context) {
	var complaint models.Complaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if complaint.AuctionID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction ID is required"})
		return
	}
	if complaint.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	// Check if auction exists
	var auction models.Auction
	if err := config.DB.First(&auction, complaint.AuctionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Auction not found"})
		return
	}

	// Set default status if not provided
	if complaint.Status == "" {
		complaint.Status = "pending"
	}

	// Validate status
	validStatuses := map[string]bool{"pending": true, "resolved": true, "rejected": true}
	if !validStatuses[complaint.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Must be one of: pending, resolved, rejected"})
		return
	}

	if err := config.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, complaint)
}

// GetComplaints retrieves all complaints
func GetComplaints(c *gin.Context) {
	var complaints []models.Complaint
	if err := config.DB.Find(&complaints).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, complaints)
}

// UpdateComplaintStatus updates the status of a complaint
func UpdateComplaintStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid complaint ID"})
		return
	}

	var complaint models.Complaint
	if err := config.DB.First(&complaint, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}

	var updateData struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate status
	validStatuses := map[string]bool{"pending": true, "resolved": true, "rejected": true}
	if !validStatuses[updateData.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Must be one of: pending, resolved, rejected"})
		return
	}

	// Update status
	complaint.Status = updateData.Status
	if err := config.DB.Save(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, complaint)
}
