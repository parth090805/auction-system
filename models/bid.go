package models

import "time"

type Bid struct {
    BidID   uint      `gorm:"primaryKey" json:"bid_id"`
    AuctionID uint    `json:"auction_id"`
    UserID  uint      `json:"user_id"`
    Amount  float64   `json:"amount"`
    BidTime time.Time `json:"bid_time"`
}
