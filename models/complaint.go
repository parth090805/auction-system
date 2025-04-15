package models

type Complaint struct {
    ComplaintID uint   `gorm:"primaryKey" json:"complaint_id"`
    AuctionID   uint   `json:"auction_id"`
    Description string `json:"description"`
    Status      string `json:"status"`
}
