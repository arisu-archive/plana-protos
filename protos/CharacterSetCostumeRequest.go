package protos

type CharacterSetCostumeRequest struct {
	RequestPacket
	CharacterUniqueId int64 `json:",omitempty,omitzero"`
	CostumeIdToSet *int64 `json:",omitempty,omitzero"`
}
