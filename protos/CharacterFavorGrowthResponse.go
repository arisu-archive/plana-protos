package protos

type CharacterFavorGrowthResponse struct {
	ResponsePacket
	CharacterDB                  *CharacterDB `json:",omitempty,omitzero"`
	ConsumeStackableItemDBResult []*ItemDB
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
}
