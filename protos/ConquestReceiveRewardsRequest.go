package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestReceiveRewardsRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Difficulty flatdata.StageDifficulty `json:",omitempty,omitzero"`
	Step int32 `json:",omitempty,omitzero"`
}
