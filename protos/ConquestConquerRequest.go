package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestConquerRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Difficulty flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId int64 `json:",omitempty,omitzero"`
	TileRewardId int64 `json:",omitempty,omitzero"`
}
