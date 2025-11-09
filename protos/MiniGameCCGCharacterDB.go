package protos

type MiniGameCCGCharacterDB struct {
	SlotIndex   int32 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	Health      int32 `json:",omitempty,omitzero"`
}
