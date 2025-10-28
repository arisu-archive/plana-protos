package protos

import (
	"time"
)

type GuideMissionSeasonDB struct {
	SeasonId int64 `json:",omitempty,omitzero"`
	LoginCount int64 `json:",omitempty,omitzero"`
	StartDate time.Time `json:",omitempty,omitzero"`
	LoginDate time.Time `json:",omitempty,omitzero"`
	IsComplete bool `json:",omitempty,omitzero"`
	IsFinalMissionComplete bool `json:",omitempty,omitzero"`
	CollectionItemReceiveDate *time.Time `json:",omitempty,omitzero"`
}
