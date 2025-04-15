package models

type User struct {
    UserID   uint   `gorm:"primaryKey" json:"user_id"`
    Username string `json:"username"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
