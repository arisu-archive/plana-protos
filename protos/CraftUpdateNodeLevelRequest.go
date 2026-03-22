package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CraftUpdateNodeLevelRequest struct {
	RequestPacket
	ConsumeRequestDB  ConsumeRequestDB
	ConsumeGoldAmount int64                  `json:",omitempty,omitzero"`
	SlotId            int64                  `json:",omitempty,omitzero"`
	CraftNodeType     flatdata.CraftNodeTier `json:",omitempty,omitzero"`
}
