package protos

type CharacterSetFavoritesResponse struct {
	ResponsePacket
	CharacterDBs []*CharacterDB
}
