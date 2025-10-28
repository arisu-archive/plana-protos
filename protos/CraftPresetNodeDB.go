package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CraftPresetNodeDB struct {
	NodeTier flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	IsActivated bool `json:",omitempty,omitzero"`
	PriortyNodeId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
