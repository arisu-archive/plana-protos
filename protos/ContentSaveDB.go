package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type ContentSaveDB struct {
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	AccountServerId int64 `json:",omitempty,omitzero"`
	CreateTime time.Time `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	LastEnterStageEchelonNumber int64 `json:",omitempty,omitzero"`
	StageEntranceFee []ParcelInfo `json:",omitempty,omitzero"`
	EnemyKillCountByUniqueId map[int64]int64 `json:",omitempty,omitzero"`
	TacticClearTimeMscSum int64 `json:",omitempty,omitzero"`
	AccountLevelWhenCreateDB int64 `json:",omitempty,omitzero"`
	BIEchelon string `json:",omitempty,omitzero"`
	BIEchelon1 string `json:",omitempty,omitzero"`
	BIEchelon2 string `json:",omitempty,omitzero"`
	BIEchelon3 string `json:",omitempty,omitzero"`
	BIEchelon4 string `json:",omitempty,omitzero"`
}
