package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CampaignChapterClearRewardRequest struct {
	RequestPacket
	CampaignChapterUniqueId int64                    `json:",omitempty,omitzero"`
	StageDifficulty         flatdata.StageDifficulty `json:",omitempty,omitzero"`
}
