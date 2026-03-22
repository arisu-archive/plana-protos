package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB       ConquestInfoDB                  `json:",omitempty,omitzero"`
	ConquestedTileDBs    []ConquestTileDB                `json:",omitempty,omitzero"`
	DifficultyToStepDict *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	IsFirstEnter         bool                            `json:",omitempty,omitzero"`
}
