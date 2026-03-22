package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestGetInfoResponse struct {
	ResponsePacket
	ConquestInfoDB           ConquestInfoDB                  `json:",omitempty,omitzero"`
	ConquestedTileDBs        []ConquestTileDB                `json:",omitempty,omitzero"`
	ConquestObjectDBsWrapper []ConquestEventObjectDB         `json:",omitempty,omitzero"`
	ConquestEchelonDBs       []ConquestEchelonDB             `json:",omitempty,omitzero"`
	DifficultyToStepDict     *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	IsFirstEnter             bool                            `json:",omitempty,omitzero"`
	DisplayInfos             []ConquestDisplayInfo           `json:",omitempty,omitzero"`
}
