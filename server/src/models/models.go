package models

import "time"

type Contacts struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	FirstName string         `gorm:"size:100;not null"`
	LastName  string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;uniqueIndex;not null"`
	Phone     string         `gorm:"size:20;uniqueIndex;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoCreateTime"`
}