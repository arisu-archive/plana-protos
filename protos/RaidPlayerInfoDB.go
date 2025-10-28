package protos

import (
	"time"
)

type RaidPlayerInfoDB struct {
	RaidServerId int64 `json:",omitempty,omitzero"`
	AccountId int64 `json:",omitempty,omitzero"`
	JoinDate time.Time `json:",omitempty,omitzero"`
	DamageAmount int64 `json:",omitempty,omitzero"`
	RaidEndRewardFlag int32 `json:",omitempty,omitzero"`
	RaidPlayCount int32 `json:",omitempty,omitzero"`
	Nickname string `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	CostumeId int64 `json:",omitempty,omitzero"`
	AccountLevel *int64 `json:",omitempty,omitzero"`
}
