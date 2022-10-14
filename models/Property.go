package models

import (
	"errors"

	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	CheckInTime  string    `gorm:"size:255;not null;default:current_timestamp" json:"checkin_time"`
	Guests       uint32    `gorm:"size:255;not null;" json:"guests"`
	Beds         uint32    `gorm:"size:255;not null;" json:"beds"`
	BedRooms     uint32    `gorm:"size:255;not null;" json:"bedrooms"`
	Bathrooms    uint32    `gorm:"size:255;not null;" json:"bathrooms"`
	Description  string    `gorm:"type:text;" json:"desc"`
	Location     string    `gorm:"size:255;not null;" json:"location"`
	Images       string    `gorm:"type:text;" json:"images"`
	Owner        User      `json:"owner"`
	OwnerID      uint32    `sql:"type:int; REFERENCES users(id)" json:"owner_id"`
	HouseRules   HouseRule `gorm:"house_rules"`
	HouseRulesID uint32    `gorm:"type:int REFERENCES house_rules(id)" json:"rule_id"`
}

func (p *Property) Validate() error {

	if p.Guests == 0 {
		return errors.New("check in time is required")
	}

	return nil
}
