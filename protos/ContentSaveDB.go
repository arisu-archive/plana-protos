package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ContentSaveDB struct {
	ContentType                 flatdata.ContentType
	AccountServerId             int64  `json:",omitempty,omitzero"`
	CreateTime                  MxTime `json:",omitempty,omitzero"`
	StageUniqueId               int64  `json:",omitempty,omitzero"`
	LastEnterStageEchelonNumber int64  `json:",omitempty,omitzero"`
	StageEntranceFee            []ParcelInfo
	EnemyKillCountByUniqueId    *mapx.OrderedMap[int64, int64]
	TacticClearTimeMscSum       int64 `json:",omitempty,omitzero"`
	AccountLevelWhenCreateDB    int64 `json:",omitempty,omitzero"`
	BIEchelon                   string
	BIEchelon1                  string
	BIEchelon2                  string
	BIEchelon3                  string
	BIEchelon4                  string
}
