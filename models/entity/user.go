package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint64         `gorm:"column:id;primaryKey" json:"id"`
	UserName   string         `gorm:"column:user_name;not null" json:"user_name"`
	Password   string         `gorm:"column:password;not null" json:"password"`
	Email      string         `gorm:"column:email;default:''" json:"email"`
	Phone      string         `gorm:"column:phone;default:''" json:"phone"`
	Sex        int            `gorm:"column:sex;default:1" json:"sex"`
	Age        int            `gorm:"column:age;default:18" json:"age"`
	Address    string         `gorm:"column:address;not null" json:"address"`
	CreateTime time.Time      `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime time.Time      `gorm:"column:update_time;not null" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"deleted_at"`
}
