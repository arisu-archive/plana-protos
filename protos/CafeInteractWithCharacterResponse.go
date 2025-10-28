package protos

type CafeInteractWithCharacterResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDB CafeDB `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
