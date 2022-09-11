package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Owner       User           `gorm:"foreignKey:id;references:ID" json:"owner_id,omitempty"`
	CheckInTime time.Time      `gorm:"size:255;not null;default:current_timestamp" json:"checkin_time"`
	Guests      uint32         `gorm:"size:255;not null;" json:"guests"`
	Beds        uint32         `gorm:"size:255;not null;" json:"beds"`
	BedRooms    uint32         `gorm:"size:255;not null;" json:"bedrooms"`
	Bathrooms   uint32         `gorm:"size:255;not null;" json:"bathrooms"`
	Amenities   pq.StringArray `gorm:"type:text[];not null;" json:"amenities"`
	Description string         `gorm:"type:text;" json:"description"`
	Location    string         `gorm:"size:255;not null;" json:"location"`
	HouseRules  pq.StringArray `gorm:"type:text[];" json:"houseRules"`
	Images      pq.StringArray `gorm:"type:text[];" json:"images"`
}
