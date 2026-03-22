package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB           ConquestInfoDB
	ConquestedTileDBs        []ConquestTileDB
	ConquestObjectDBsWrapper []ConquestEventObjectDB
	ConquestEchelonDBs       []ConquestEchelonDB
	DifficultyToStepDict     *mapx.OrderedMap[string, int32]
	IsFirstEnter             bool `json:",omitempty,omitzero"`
	DisplayInfos             []ConquestDisplayInfo
}
