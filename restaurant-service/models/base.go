package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseID struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type BaseUID struct {
	ID string `json:"id" gorm:"primarykey"`
}

type BaseDateTimeMeta struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
