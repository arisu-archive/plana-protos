package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestMainStorySummary struct {
	EventContentId          int64                    `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty `json:",omitempty,omitzero"`
	ConquestStepSummaryDict *mapx.OrderedMap[int32, *ConquestMainStoryStepSummary]
}
