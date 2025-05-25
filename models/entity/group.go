package entity

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID           uint64         `gorm:"column:id;primaryKey" json:"id"`
	GroupName    string         `gorm:"column:group_name;not null" json:"group_name"`
	CreateTime   time.Time      `gorm:"column:create_time;not null" json:"create_time"`
	RememberNum  int            `gorm:"column:remember_num;not null" json:"remember_num"`
	Describe     string         `gorm:"column:describe;not null" json:"describe"`
	CreateUserID uint64         `gorm:"column:create_user_id;not null" json:"create_user_id"`
	UpdateUserID uint64         `gorm:"column:update_user_id;not null" json:"update_user_id"`
	DeleteUserID uint64         `gorm:"column:delete_user_id;not null" json:"delete_user_id"`
	UpdateTime   time.Time      `gorm:"column:update_time;not null" json:"update_time"`
	DeleteTime   gorm.DeletedAt `gorm:"column:delete_time" json:"deleted_at"`
}
