package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventContentMainStageSaveDB struct {
	CampaignMainStageSaveDB
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	SelectedBuffDict map[int64]int64 `json:",omitempty,omitzero"`
	CurrentAppearedBuffGroupId int64 `json:",omitempty,omitzero"`
}
