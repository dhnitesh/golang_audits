package models

import "time"

type GeoData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Data struct {
	Created   time.Time `json:"created"`
	User      string    `json:"user"`
	RawUserID int       `json:"raw_user_id"`
	Username  string    `json:"username"`
	IP        string    `json:"ip"`
	Host      string    `json:"host"`
	Path      string    `json:"path"`
	UserAgent string    `json:"user_agent"`
	UUID      string    `json:"uuid"`
	ExitTime  time.Time `json:"exit_time"`
	GeoData   GeoData   `json:"geo_data"`
}