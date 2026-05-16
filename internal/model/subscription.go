package model

import "time"

type Subscription struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SubscriptionResponse struct {
	ID          uint   `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date,omitempty"`
}
