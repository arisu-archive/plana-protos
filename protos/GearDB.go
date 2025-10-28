package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type GearDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	Exp int64 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	SlotIndex int64 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
