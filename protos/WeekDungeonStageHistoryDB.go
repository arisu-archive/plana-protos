package protos

import (
	"github.com/arisu-archive/mapx"
)

type WeekDungeonStageHistoryDB struct {
	AccountServerId int64 `json:",omitempty,omitzero"`
	StageUniqueId   int64 `json:",omitempty,omitzero"`
	StarGoalRecord  *mapx.OrderedMap[string, int64]
}
