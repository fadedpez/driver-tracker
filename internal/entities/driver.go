package entities

import (
	"time"
)

type Driver struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Number      string `json:"driver_number"`
	Nationality string `json:"driver_nationality"`
	dob         *time.Time
	ID          string `json:"id"`
}
