package models

import "time"

type Transaction struct {
    TransactionID uint      `gorm:"primaryKey" json:"transaction_id"`
    AuctionID     uint      `json:"auction_id"`
    Amount        float64   `json:"amount"`
    Date          time.Time `json:"date"`
}
