package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TimeAttackDungeonBattleHistoryDB struct {
	DungeonType         flatdata.TimeAttackDungeonType `json:",omitempty,omitzero"`
	GeasId              int64                          `json:",omitempty,omitzero"`
	DefaultPoint        int64                          `json:",omitempty,omitzero"`
	ClearTimePoint      int64                          `json:",omitempty,omitzero"`
	EndFrame            int64                          `json:",omitempty,omitzero"`
	MainCharacterDBs    []TimeAttackDungeonCharacterDB
	SupportCharacterDBs []TimeAttackDungeonCharacterDB
}
