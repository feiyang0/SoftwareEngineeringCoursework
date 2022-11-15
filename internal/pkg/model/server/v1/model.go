package v1

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `json:"id,omitempty"        gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deletedAt;index:idx_deletedAt"`
}
