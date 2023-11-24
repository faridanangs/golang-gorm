package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int            `gorm:"primary_key;column:id;autoIncrement"`
	UserId      string         `gorm:"column:user_id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *Todo) TableName() string {
	return "todos"
}
