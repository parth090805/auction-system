package models

import "time"

type Auction struct {
    AuctionID     uint      `gorm:"primaryKey" json:"auction_id"`
    ItemID        uint      `json:"item_id"`
    StartingPrice float64   `json:"starting_price"`
    StartTime     time.Time `json:"start_time"`
    EndTime       time.Time `json:"end_time"`
}
