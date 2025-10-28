package protos

type CharacterSetCostumeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SetCostumeDB CostumeDB `json:",omitempty,omitzero"`
	UnsetCostumeDB CostumeDB `json:",omitempty,omitzero"`
}
