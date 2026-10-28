package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonSuggestionRequest struct {
	RequestPacket
	EchelonType                flatdata.EchelonType          `json:",omitempty,omitzero"`
	EchelonNumber              int64                         `json:",omitempty,omitzero"`
	ExtensionType              flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	ContentType                flatdata.ContentType
	StageId                    int64 `json:",omitempty,omitzero"`
	IsStrategySkip             bool  `json:",omitempty,omitzero"`
	IsPreset                   bool  `json:",omitempty,omitzero"`
	AllowSamePersonality       bool  `json:",omitempty,omitzero"`
	ExcludedCharacterServerIds []int64
}
