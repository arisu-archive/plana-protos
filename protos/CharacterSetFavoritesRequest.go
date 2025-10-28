package protos

type CharacterSetFavoritesRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ActivateByServerIds map[int64]bool `json:",omitempty,omitzero"`
}
