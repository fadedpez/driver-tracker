package entities

import (
	"time"

	"github.com/fadedpez/driver-tracker/protos"
)

//The goal here is to pull in all the drivers that participated in Qualifying. It is possible for this list to be different than the Grand Prix drivers (not common).
//Participants (drivers) belong to a qualifying session.
//Teams belong to a quali session too? Maybe! I dunno!

type Qualifying struct {
	GrandPrixName  string                          `json:"qualifying_grand_prix_name"`
	GrandPrixTrack string                          `json:"qualifying_grand_prix_track"`
	Date           *time.Time                      `json:"qualifying_date"`
	Participants   *[]protos.QualifyingParticipant `json:"qualifying_participants"`
	ID             string                          `json:"qualifying_id"`
}

//TODO: Need to see if this does what I intend.
