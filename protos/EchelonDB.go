package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonDB struct {
	AccountServerId               int64                         `json:",omitempty,omitzero"`
	EchelonType                   flatdata.EchelonType          `json:",omitempty,omitzero"`
	EchelonNumber                 int64                         `json:",omitempty,omitzero"`
	ExtensionType                 flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	LeaderServerId                int64                         `json:",omitempty,omitzero"`
	MainSlotServerIds             []int64
	SupportSlotServerIds          []int64
	TSSInteractionServerId        int64                       `json:",omitempty,omitzero"`
	UsingFlag                     EchelonDB_EchelonStatusFlag `json:",omitempty,omitzero"`
	SkillCardMulliganCharacterIds []int64
	CombatStyleIndex              []int32
}
