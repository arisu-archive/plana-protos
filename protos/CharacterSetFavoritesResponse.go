package protos

type CharacterSetFavoritesResponse struct {
	ResponsePacket
	CharacterDBs []CharacterDB `json:",omitempty,omitzero"`
}
