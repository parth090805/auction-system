package models

import "time"

type Notification struct {
    NotificationID uint      `gorm:"primaryKey" json:"notification_id"`
    UserID         uint      `json:"user_id"`
    Message        string    `json:"message"`
    Timestamp      time.Time `json:"timestamp"`
}
