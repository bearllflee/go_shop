package global

import (
	"gorm.io/gorm"
	"time"
)

type GSModel struct {
	ID        uint64         `gorm:"primaryKey" json:"ID"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
