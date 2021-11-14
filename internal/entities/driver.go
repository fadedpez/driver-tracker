package entities

import (
	"time"
)

type Driver struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	DriverNumber      string `json:"driver_number"`
	DriverNationality string `json:"driver_nationality"`
	dob               *time.Time
}
