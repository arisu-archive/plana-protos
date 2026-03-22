package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ArenaTeamSettingDB struct {
	EchelonType               flatdata.EchelonType `json:",omitempty,omitzero"`
	LeaderCharacterId         int64                `json:",omitempty,omitzero"`
	TSSInteractionCharacterId int64                `json:",omitempty,omitzero"`
	MainCharacters            []ArenaCharacterDB
	SupportCharacters         []ArenaCharacterDB
	TSSCharacterDB            ArenaCharacterDB
	MapId                     int64 `json:",omitempty,omitzero"`
}
