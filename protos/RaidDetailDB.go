package protos

import (
	"time"
)

type RaidDetailDB struct {
	RaidUniqueId int64 `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	DamageTable []RaidPlayerInfoDB `json:",omitempty,omitzero"`
}
