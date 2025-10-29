package protos

type CharacterGearUnlockRequest struct {
	RequestPacket
	CharacterServerId int64 `json:",omitempty,omitzero"`
	SlotIndex int32 `json:",omitempty,omitzero"`
}
