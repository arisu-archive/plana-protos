package protos

import (
	"time"
)

type WorldRaidClearHistoryDB struct {
	SeasonId int64 `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	RewardReceiveDate time.Time `json:",omitempty,omitzero"`
}
