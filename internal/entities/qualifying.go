package entities

import (
	"time"

	"github.com/fadedpez/driver-tracker/protos"
)

type Qualifying struct {
	QualifyingGrandPrixName  string                          `json:"qualifying_grand_prix_name"`
	QualifyingGrandPrixTrack string                          `json:"qualifying_grand_prix_track"`
	QualifyingDate           *time.Time                      `json:"qualifying_date"`
	QualifyingParticipants   *[]protos.QualifyingParticipant `json:"qualifying_participants"`
	QualifyingID             string                          `json:"qualifying_id"`
}
