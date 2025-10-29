package protos

type CharacterFavorGrowthResponse struct {
	ResponsePacket
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	ConsumeStackableItemDBResult []ItemDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
