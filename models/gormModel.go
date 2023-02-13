package models

import "time"

type GormModel struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Created_At time.Time `json:"created_at,omitempty"`
	Updated_At time.Time `json:"updated_at,omitempty"`
}
