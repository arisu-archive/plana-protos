package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestConquerDeployEchelonRequest struct {
	RequestPacket
	EventContentId    int64                    `json:",omitempty,omitzero"`
	Difficulty        flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId      int64                    `json:",omitempty,omitzero"`
	EchelonDB         EchelonDB
	ClanAssistUseInfo ClanAssistUseInfo
}
