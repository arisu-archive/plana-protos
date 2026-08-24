package protos

type CharacterAllocateEngraveNodeResponse struct {
	ResponsePacket
	CharacterDB    *CharacterDB    `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
