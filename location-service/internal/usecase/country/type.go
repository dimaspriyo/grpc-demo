package country

import "time"

type CountryResponse struct {
	Id         int64     `json:"id,omitempty"`
	Country    string    `json:"country,omitempty"`
	LastUpdate time.Time `json:"last_update,omitempty"`
}
