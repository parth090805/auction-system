package controllers

import (
	"auction-system/config"
	"auction-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTransaction handles the creation of a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction data: " + err.Error()})
		return
	}

	// Validate required fields
	if transaction.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be positive"})
		return
	}

	if transaction.WalletID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet ID is required"})
		return
	}

	// Set default status if not provided
	if transaction.Status == "" {
		transaction.Status = "pending"
	}

	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransactions retrieves all transactions with optional filtering
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	// Add optional query parameters for filtering
	query := config.DB
	if walletID := c.Query("wallet_id"); walletID != "" {
		query = query.Where("wallet_id = ?", walletID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransactionByID retrieves a specific transaction by its ID
func GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
