package models

import "time"

type Wallet struct {
    WalletID    uint      `gorm:"primaryKey" json:"wallet_id"`
    UserID      uint      `json:"user_id"`
    Balance     float64   `json:"balance"`
    LastUpdated time.Time `json:"last_updated"`
}
