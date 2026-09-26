package domain

import (
	"time"
)

type User struct {
	ID               int64     `json:"id" gorm:"type:uuid;primaryKey;not null"`
	CityName         string    `json:"cityName"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Timezone         string    `json:"timezone"`
	Units            string    `json:"units"`
	NotificationTime time.Time `json:"notification_time"`
	CreatedAt        time.Time `json:"createdAt" gorm:"default:now()"`
}
