package models

import (
	"gorm.io/gorm"
)

type HouseRule struct {
	gorm.Model
	Rule string `gorm:"type:text;" json:"rule"`
}
