package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type SystemAssistInfo struct {
	AssistType       SystemAssistType `json:",omitempty,omitzero"`
	ContentType      flatdata.ContentType
	StageId          int64 `json:",omitempty,omitzero"`
	AssistCharacters []*SystemAssistCharacter
}
