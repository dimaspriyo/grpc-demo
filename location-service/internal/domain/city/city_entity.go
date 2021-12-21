package city

import "time"

type City struct {
	CityId     int    `gorm:"primaryKey"`
	City       string `gorm:"size:255"`
	CountryId  int
	LastUpdate time.Time
}
