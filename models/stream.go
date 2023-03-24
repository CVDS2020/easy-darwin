package models

type Stream struct {
	URL               string `gorm:"type:varchar(255);primary_key;unique"`
	CustomPath        string `gorm:"type:varchar(255)"`
	IdleTimeout       int
	HeartbeatInterval int
}
