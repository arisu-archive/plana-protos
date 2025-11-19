package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CraftPresetNodeDB struct {
	NodeTier        flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	IsActivated     bool                   `json:",omitempty,omitzero"`
	PriorityNodeIds []int64                `json:",omitempty,omitzero"`
	CostParcels     []ParcelInfoImmutable  `json:",omitempty,omitzero"`
}
