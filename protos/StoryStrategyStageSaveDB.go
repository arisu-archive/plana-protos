package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type StoryStrategyStageSaveDB struct {
	CampaignMainStageSaveDB
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
}
