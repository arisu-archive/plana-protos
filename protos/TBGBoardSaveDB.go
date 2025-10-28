package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGBoardSaveDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Round int32 `json:",omitempty,omitzero"`
	ThemaIndex int32 `json:",omitempty,omitzero"`
	CurrentThemaMapType flatdata.TBGThemaType `json:",omitempty,omitzero"`
	MainMap TBGHexaMapDB `json:",omitempty,omitzero"`
	HiddenMap TBGHexaMapDB `json:",omitempty,omitzero"`
	Player TBGPlayerDB `json:",omitempty,omitzero"`
	Encounter TBGEncounterDB `json:",omitempty,omitzero"`
	BestClearRecord map[int32]TBGThemaClearRecord `json:",omitempty,omitzero"`
	HiddenTreasureRecord []int32 `json:",omitempty,omitzero"`
	HiddenPotalOpenConditionRecord []int32 `json:",omitempty,omitzero"`
}
