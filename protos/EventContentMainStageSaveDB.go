package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentMainStageSaveDB struct {
	CampaignMainStageSaveDB
	SelectedBuffDict           *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	CurrentAppearedBuffGroupId int64                          `json:",omitempty,omitzero"`
}
