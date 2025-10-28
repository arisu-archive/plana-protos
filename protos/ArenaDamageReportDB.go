package protos

import (
	"time"
)

type ArenaDamageReportDB struct {
	ArenaBattleServerId int64 `json:",omitempty,omitzero"`
	WinnerAccountServerId int64 `json:",omitempty,omitzero"`
	AttackerUserDB ArenaUserDB `json:",omitempty,omitzero"`
	DefenderUserDB ArenaUserDB `json:",omitempty,omitzero"`
	BattleEndTime time.Time `json:",omitempty,omitzero"`
	AttackerDamageReport map[int64]int64 `json:",omitempty,omitzero"`
	DefenderDamageReport map[int64]int64 `json:",omitempty,omitzero"`
}
