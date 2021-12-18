package entities

import (
	"time"

	"github.com/fadedpez/driver-tracker/protos"
)

type GrandPrix struct {
	GrandPrixName         string                         `json:"grand_prix_name"`
	GrandPrixTrack        string                         `json:"grand_prix_track"`
	GrandPrixDate         *time.Time                     `json:"grand_prix_date"`
	GrandPrixParticipants *[]protos.GrandPrixParticipant `json:"grand_prix_participants"`
	GrandPrixID           string                         `json:"grand_prix_id"`
}
