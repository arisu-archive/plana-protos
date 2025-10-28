package protos

type CharacterGearUnlockRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterServerId int64 `json:",omitempty,omitzero"`
	SlotIndex int32 `json:",omitempty,omitzero"`
}
