package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB       *ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestedTileDBs    []*ConquestTileDB
	DifficultyToStepDict *mapx.OrderedMap[string, int32]
	IsFirstEnter         bool `json:",omitempty,omitzero"`
}
