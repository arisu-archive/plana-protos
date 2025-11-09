package protos

type CharacterGearUnlockResponse struct {
	ResponsePacket
	GearDB      GearDB      `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
}
