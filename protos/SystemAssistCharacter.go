package protos

type SystemAssistCharacter struct {
	CharacterId      int64           `json:",omitempty,omitzero"`
	EchelonSlotType  EchelonSlotType `json:",omitempty,omitzero"`
	EchelonSlotIndex int32           `json:",omitempty,omitzero"`
	CombatStyleIndex int32           `json:",omitempty,omitzero"`
	IsMulligan       bool            `json:",omitempty,omitzero"`
}
