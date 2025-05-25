package entity

import (
	"gorm.io/gorm"
	"time"
)

type GroupUser struct {
	ID         uint64         `gorm:"column:id;primaryKey" json:"id"`
	GroupID    uint64         `gorm:"column:group_id;not null" json:"group_id"`
	UserID     uint64         `gorm:"column:user_id;not null" json:"user_id"`
	Type       int8           `gorm:"column:type;default:2" json:"type"` // 成员类型: 0:群主, 1:管理员, 2:普通成员
	CreateTime time.Time      `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime time.Time      `gorm:"column:update_time;not null" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"deleted_at"`
}
