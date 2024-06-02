package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"size:64;index;unique;not null"`
	Password string `gorm:"size:64;not null"`
	Role     string `gorm:"type:enum('admin', 'sales', 'agents', 'guest');default:'guest';not null"`
}
