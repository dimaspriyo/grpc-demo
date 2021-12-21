package country

import "time"

type Country struct {
	CountryId  int64  `gorm:"primaryKey"`
	Country    string `gorm:"type:varchar(50)"`
	LastUpdate time.Time
}

func (Country) TableName() string {
	return "country"
}
