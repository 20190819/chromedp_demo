package model

import (
	"database/sql"
	"github.com/uniplaces/carbon"
	"time"
)

type Pk struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type TimeAt struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

type CreatedAt struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
}

type UpdatedAt struct {
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

type DeletedAt struct {
	DeletedAt sql.NullTime `gorm:"index"`
}

func SerializeDate(t time.Time) string {
	return carbon.NewCarbon(t).DateTimeString()
}