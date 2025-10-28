package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonDB struct {
	AccountServerId int64 `json:",omitempty,omitzero"`
	EchelonType flatdata.EchelonType `json:",omitempty,omitzero"`
	EchelonNumber int64 `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	LeaderServerId int64 `json:",omitempty,omitzero"`
	MainSlotServerIds []int64 `json:",omitempty,omitzero"`
	SupportSlotServerIds []int64 `json:",omitempty,omitzero"`
	TSSInteractionServerId int64 `json:",omitempty,omitzero"`
	UsingFlag EchelonDB_EchelonStatusFlag `json:",omitempty,omitzero"`
	SkillCardMulliganCharacterIds []int64 `json:",omitempty,omitzero"`
	CombatStyleIndex []int32 `json:",omitempty,omitzero"`
}
