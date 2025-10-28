package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CharacterDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	StarGrade int32 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	Exp int64 `json:",omitempty,omitzero"`
	FavorRank int32 `json:",omitempty,omitzero"`
	FavorExp int64 `json:",omitempty,omitzero"`
	PublicSkillLevel int32 `json:",omitempty,omitzero"`
	ExSkillLevel int32 `json:",omitempty,omitzero"`
	PassiveSkillLevel int32 `json:",omitempty,omitzero"`
	ExtraPassiveSkillLevel int32 `json:",omitempty,omitzero"`
	LeaderSkillLevel int32 `json:",omitempty,omitzero"`
	IsFavorite bool `json:",omitempty,omitzero"`
	EquipmentServerIds []int64 `json:",omitempty,omitzero"`
	PotentialStats map[int32]int32 `json:",omitempty,omitzero"`
}
