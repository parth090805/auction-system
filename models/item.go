package models

type Item struct {
    ItemID     uint   `gorm:"primaryKey" json:"item_id"`
    OwnerID    uint   `json:"owner_id"`
    SellerID   uint   `json:"seller_id"`
    Description string `json:"description"`
    Images      string `json:"images"`
    CurrentBid  float64 `json:"current_bid"`
    SoldFor     float64 `json:"sold_for"`
}
