package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	HostName    User
	CheckInTime time.Time      `gorm:"size:255;not null;default:flexible" json:"checkin_time"`
	Guests      uint32         `gorm:"size:255;not null;" json:"guests"`
	Beds        uint32         `gorm:"size:255;not null;" json:"beds"`
	BedRooms    uint32         `gorm:"size:255;not null;" json:"bedrooms"`
	Bathrooms   uint32         `gorm:"size:255;not null;" json:"bathrooms"`
	Amenities   pq.StringArray `gorm:"size:255;not null;" json:"amenities"`
	Description string         `gorm:"type:text;not null;" json:"description"`
	Location    string         `gorm:"size:255;not null;" json:"location"`
	HouseRules  pq.StringArray `gorm:"size:255;" json:"houseRules"`
}
