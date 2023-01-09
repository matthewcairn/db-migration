package model

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Pagination struct {
	Page     int
	PageSize int
}

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

// MsMetadata describes the structure.
type BaseModel struct {
	ID        uuid.UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatorID *uuid.UUID      `json:"creator_id,omitempty"`
	UpdaterID *uuid.UUID      `json:"updater_id,omitempty"`
	CreatedAt time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type StandardData struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt int       `json:"created_at"`
	Data      string    `json:"data"`
	UserID    string    `json:"user_id"`
}

type DataConsumer struct {
	Before interface{} `json:"before"`
	After  Data        `json:"after"`
}

type Data struct {
	Id uuid.UUID `json:"id"`
}
