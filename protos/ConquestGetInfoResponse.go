package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestedTileDBs []ConquestTileDB `json:",omitempty,omitzero"`
	ConquestObjectDBsWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	ConquestEchelonDBs []ConquestEchelonDB `json:",omitempty,omitzero"`
	DifficultyToStepDict map[flatdata.StageDifficulty]int32 `json:",omitempty,omitzero"`
	IsFirstEnter bool `json:",omitempty,omitzero"`
	DisplayInfos []ConquestDisplayInfo `json:",omitempty,omitzero"`
}
