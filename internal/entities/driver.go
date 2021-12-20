package entities

import (
	"time"
)

type Driver struct {
	FirstName   string `dynamodbav:"first_name"`
	LastName    string `dynamodbav:"last_name"`
	Number      string `dynamodbav:"driver_number"`
	Nationality string `dynamodbav:"driver_nationality"`
	dob         *time.Time
	ID          string `dynamodbav:"id"`
}
