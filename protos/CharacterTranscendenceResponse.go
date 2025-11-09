package protos

type CharacterTranscendenceResponse struct {
	ResponsePacket
	CharacterDB    CharacterDB    `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
