package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestStageSaveDB struct {
	ContentSaveDB
	ConquestEventObjectDBId *int64                    `json:",omitempty,omitzero"`
	EventContentId          int64                     `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty  `json:",omitempty,omitzero"`
	TileUniqueId            int64                     `json:",omitempty,omitzero"`
	TilePresetId            int64                     `json:",omitempty,omitzero"`
	ConquestTileType        flatdata.ConquestTileType `json:",omitempty,omitzero"`
	UseManageEchelon        bool                      `json:",omitempty,omitzero"`
	AssistCharacterDB       AssistCharacterDB         `json:",omitempty,omitzero"`
	EchelonSlotType         int32                     `json:",omitempty,omitzero"`
	EchelonSlotIndex        int32                     `json:",omitempty,omitzero"`
}
