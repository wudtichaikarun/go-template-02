package model

import (
	"strconv"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (m *Model) IDtoString() string {
	return strconv.Itoa(int(m.ID))
}
