package protos

type CafeInteractWithCharacterResponse struct {
	ResponsePacket
	CafeDB CafeDB `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
