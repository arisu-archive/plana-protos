package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ArenaTeamSettingDB struct {
	EchelonType flatdata.EchelonType `json:",omitempty,omitzero"`
	LeaderCharacterId int64 `json:",omitempty,omitzero"`
	TSSInteractionCharacterId int64 `json:",omitempty,omitzero"`
	MainCharacters []ArenaCharacterDB `json:",omitempty,omitzero"`
	SupportCharacters []ArenaCharacterDB `json:",omitempty,omitzero"`
	TSSCharacterDB ArenaCharacterDB `json:",omitempty,omitzero"`
	MapId int64 `json:",omitempty,omitzero"`
}
