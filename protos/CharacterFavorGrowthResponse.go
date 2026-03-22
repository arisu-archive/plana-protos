package protos

type CharacterFavorGrowthResponse struct {
	ResponsePacket
	CharacterDB                  CharacterDB
	ConsumeStackableItemDBResult []ItemDB
	ParcelResultDB               ParcelResultDB
}
