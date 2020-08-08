package model

import "time"

type BaseModel struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	//Deleted     bool      `json:"deleted"`
	//DeletedTime time.Time `json:"deleted_time"`
}
