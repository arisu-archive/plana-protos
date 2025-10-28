package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type CampaignChapterClearRewardHistoryDB struct {
	ChapterUniqueId int64 `json:",omitempty,omitzero"`
	RewardType flatdata.StageDifficulty `json:",omitempty,omitzero"`
	ReceiveDate time.Time `json:",omitempty,omitzero"`
}
