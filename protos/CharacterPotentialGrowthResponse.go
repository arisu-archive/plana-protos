package protos

type CharacterPotentialGrowthResponse struct {
	ResponsePacket
	CharacterDB    CharacterDB    `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
