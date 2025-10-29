package protos

type CharacterSetCostumeResponse struct {
	ResponsePacket
	SetCostumeDB CostumeDB `json:",omitempty,omitzero"`
	UnsetCostumeDB CostumeDB `json:",omitempty,omitzero"`
}
