package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestMainStoryGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB       ConquestInfoDB
	ConquestedTileDBs    []ConquestTileDB
	DifficultyToStepDict *mapx.OrderedMap[string, int32]
	IsFirstEnter         bool `json:",omitempty,omitzero"`
}
