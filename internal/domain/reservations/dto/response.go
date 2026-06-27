package dto

import (
	"time"
)

type ReservationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ZoneID       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ZoneResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type MyReservationResponse struct {
	ID           uint         `json:"id"`
	LicensePlate string       `json:"license_plate"`
	Status       string       `json:"status"`
	Zone         ZoneResponse `json:"zone"`
	CreatedAt    time.Time    `json:"created_at"`
}

type UserSummaryResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type AdminReservationResponse struct {
	ID           uint                `json:"id"`
	UserID       uint                `json:"user_id"`
	ZoneID       uint                `json:"zone_id"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	User         UserSummaryResponse `json:"user"`
	Zone         ZoneResponse        `json:"zone"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}
