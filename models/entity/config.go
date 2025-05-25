package entity

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	ID          uint64         `gorm:"column:id;primaryKey" json:"id"`
	ConfigName  string         `gorm:"column:config_name;default:''" json:"config_name"`
	ConfigCode  string         `gorm:"column:config_code;default:''" json:"config_code"`
	ConfigValue string         `gorm:"column:config_value;default:''" json:"config_value"`
	Describe    string         `gorm:"column:describe;not null" json:"describe"`
	CreateTime  time.Time      `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime  time.Time      `gorm:"column:update_time;not null" json:"update_time"`
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time" json:"deleted_at"`
}
