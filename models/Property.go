package models

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	ID          uint32         `gorm:"primary_key;auto_increment" json:"id"`
	CheckInTime string         `gorm:"size:255;not null;default:current_timestamp" json:"checkin_time"`
	Guests      uint32         `gorm:"size:255;not null;" json:"guests"`
	Beds        uint32         `gorm:"size:255;not null;" json:"beds"`
	BedRooms    uint32         `gorm:"size:255;not null;" json:"bedrooms"`
	Bathrooms   uint32         `gorm:"size:255;not null;" json:"bathrooms"`
	Description string         `gorm:"type:text;" json:"desc"`
	Location    string         `gorm:"size:255;not null;" json:"location"`
	Images      pq.StringArray `gorm:"type:varchar(64)[]" json:"images"`
	HouseRules  pq.StringArray `gorm:"type:varchar(64)[]" json:"house_rules"`
	Owner       User           `json:"user"`
	OwnerId     uint32         `sql:"type:int REFERENCES user(id)" json:"user_id"`
}

func (p *Property) Validate() error {

	if p.Guests == 0 {
		return errors.New("the number of guests is required")
	}
	if p.CheckInTime == "" {
		return errors.New("check in time is required")
	}
	if p.Beds == 0 {
		return errors.New("number of beds is required")
	}
	if p.Bathrooms == 0 {
		return errors.New("number of bathrooms is required")
	}
	if p.Description == "" {
		return errors.New("description is required")
	}
	if p.Location == "" {
		return errors.New("location is required")
	}
	if p.Images[0] == "" {
		return errors.New("images are required")
	}
	if p.HouseRules[0] == "" {
		return errors.New("house rules are required")
	}

	return nil
}
