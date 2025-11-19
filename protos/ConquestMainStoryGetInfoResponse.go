package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB       ConquestInfoDB                     `json:",omitempty,omitzero"`
	ConquestedTileDBs    []ConquestTileDB                   `json:",omitempty,omitzero"`
	DifficultyToStepDict map[flatdata.StageDifficulty]int32 `json:",omitempty,omitzero"`
	IsFirstEnter         bool                               `json:",omitempty,omitzero"`
}
