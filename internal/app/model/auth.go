package model

import "time"

type Auth struct {
	ID       int       `gorm:"type:bigint;primaryKey;autoIncrement"`
	Token    string    `gorm:"type:varchar"`
	AuthType string    `gorm:"type:varchar"`
	UserID   int       `gorm:"type:bigint;column:user_id;not null"`
	Expiry   time.Time `gorm:"column:expires_at;type:timestamp"`
}
