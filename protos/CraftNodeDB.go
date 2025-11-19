package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CraftNodeDB struct {
	NodeTier           flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	SlotSequence       int64                  `json:",omitempty,omitzero"`
	NodeId             int64                  `json:",omitempty,omitzero"`
	NodeQuality        int64                  `json:",omitempty,omitzero"`
	NodeLevel          int64                  `json:",omitempty,omitzero"`
	NodeRandomSeed     int32                  `json:",omitempty,omitzero"`
	NodeRandomSequence int32                  `json:",omitempty,omitzero"`
	LeafNodeIds        []int64                `json:",omitempty,omitzero"`
	ResultId           int64                  `json:",omitempty,omitzero"`
	CraftNodeResult    CraftNodeResult        `json:",omitempty,omitzero"`
}
