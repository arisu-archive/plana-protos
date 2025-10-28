package protos

type CharacterSetCostumeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterUniqueId int64 `json:",omitempty,omitzero"`
	CostumeIdToSet *int64 `json:",omitempty,omitzero"`
}
