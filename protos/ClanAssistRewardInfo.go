package protos

import (
	"time"
)

type ClanAssistRewardInfo struct {
	CharacterDBId int64 `json:",omitempty,omitzero"`
	DeployDate time.Time `json:",omitempty,omitzero"`
	RentCount int64 `json:",omitempty,omitzero"`
	CumultativeRewardParcels []ParcelInfo `json:",omitempty,omitzero"`
	RentRewardParcels []ParcelInfo `json:",omitempty,omitzero"`
}
