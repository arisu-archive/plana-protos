package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CampaignChapterClearRewardHistoryDB struct {
	ChapterUniqueId int64                    `json:",omitempty,omitzero"`
	RewardType      flatdata.StageDifficulty `json:",omitempty,omitzero"`
	ReceiveDate     MxTime                   `json:",omitempty,omitzero"`
}
