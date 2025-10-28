package protos

type CharacterGearUnlockResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GearDB GearDB `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
}
