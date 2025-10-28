package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestEventObjectDB struct {
	ConquestObjectDBId int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Difficulty flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId int64 `json:",omitempty,omitzero"`
	ObjectId int64 `json:",omitempty,omitzero"`
	ObjectType ConquestEventObjectType `json:",omitempty,omitzero"`
	IsAlive bool `json:",omitempty,omitzero"`
}
