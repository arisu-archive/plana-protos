package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidTeamSettingDB struct {
	AccountId                     int64                         `json:",omitempty,omitzero"`
	TryNumber                     int64                         `json:",omitempty,omitzero"`
	EchelonType                   flatdata.EchelonType          `json:",omitempty,omitzero"`
	EchelonExtensionType          flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	MainCharacterDBs              []RaidCharacterDB
	SupportCharacterDBs           []RaidCharacterDB
	SkillCardMulliganCharacterIds []int64
	TSSInteractionUniqueId        int64 `json:",omitempty,omitzero"`
	LeaderCharacterUniqueId       int64 `json:",omitempty,omitzero"`
}
