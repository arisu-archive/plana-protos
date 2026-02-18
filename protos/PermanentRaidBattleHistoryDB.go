package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type PermanentRaidBattleHistoryDB struct {
	StageId                       int64               `json:",omitempty,omitzero"`
	StartDate                     MxTime              `json:",omitempty,omitzero"`
	Status                        flatdata.RaidStatus `json:",omitempty,omitzero"`
	RaidBattleDB                  RaidBattleDB        `json:",omitempty,omitzero"`
	RaidBossDBs                   []RaidBossDB        `json:",omitempty,omitzero"`
	ParticipateCharacterServerIds []int64             `json:",omitempty,omitzero"`
	AssistUseInfo                 ClanAssistUseInfo   `json:",omitempty,omitzero"`
}
