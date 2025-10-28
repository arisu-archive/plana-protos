package protos

type CharacterSetFavoritesResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterDBs []CharacterDB `json:",omitempty,omitzero"`
}
