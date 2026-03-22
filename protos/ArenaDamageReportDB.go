package protos

import (
	"github.com/arisu-archive/mapx"
)

type ArenaDamageReportDB struct {
	ArenaBattleServerId   int64                          `json:",omitempty,omitzero"`
	WinnerAccountServerId int64                          `json:",omitempty,omitzero"`
	AttackerUserDB        ArenaUserDB                    `json:",omitempty,omitzero"`
	DefenderUserDB        ArenaUserDB                    `json:",omitempty,omitzero"`
	BattleEndTime         MxTime                         `json:",omitempty,omitzero"`
	AttackerDamageReport  *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	DefenderDamageReport  *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
}
