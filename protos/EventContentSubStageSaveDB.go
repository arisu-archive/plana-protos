package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventContentSubStageSaveDB struct {
	CampaignSubStageSaveDB
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
}
