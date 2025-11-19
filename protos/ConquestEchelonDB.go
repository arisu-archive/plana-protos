package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestEchelonDB struct {
	EventContentId          int64                    `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId            int64                    `json:",omitempty,omitzero"`
	EchelonDB               EchelonDB                `json:",omitempty,omitzero"`
	AssistCharacterUniqueId int64                    `json:",omitempty,omitzero"`
	AssistUseInfo           ClanAssistUseInfo        `json:",omitempty,omitzero"`
}
