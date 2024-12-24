package models

import (
	"gorm.io/gorm"
	"time"
)

type ActivityType struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`      // 主键ID
	TypeName  string         `gorm:"type:varchar(255);column:type_name;null"` // 类型名称
	ImageID   uint           // 类型图标ID
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// 多表查询信息存储
	Image Image `gorm:"foreignKey:ImageID" json:"image"`
}
