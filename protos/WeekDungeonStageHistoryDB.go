package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WeekDungeonStageHistoryDB struct {
	AccountServerId int64 `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	StarGoalRecord map[flatdata.StarGoalType]int64 `json:",omitempty,omitzero"`
}
