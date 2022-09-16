package models

import "time"

type History struct {
	HistoryId     uint      `gorm:"primaryKey" json:"history_id"`
	UserId        string    `json:"user_id"`
	User          User      `json:"user_data"`
	VehicleId     uint      `json:"vehicle_id"`
	Vehicle       Vehicle   `json:"vehicle_data"`
	StartRent     string    `json:"start_rent"`
	EndRent       string    `json:"end_rent"`
	Prepayment    int       `json:"prepayment"`
	PaymentStatus string    `json:"payment_status"`
	ReturnStatus  string    `json:"return_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Histories []History
