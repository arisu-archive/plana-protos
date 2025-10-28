package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CampaignTutorialStageSaveDB struct {
	ContentSaveDB
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
}
