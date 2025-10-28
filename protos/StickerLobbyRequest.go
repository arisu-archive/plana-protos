package protos

type StickerLobbyRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AcquireStickerUniqueIds []int64 `json:",omitempty,omitzero"`
}
