package protos

type CharacterSetFavoritesRequest struct {
	RequestPacket
	ActivateByServerIds map[int64]bool `json:",omitempty,omitzero"`
}
