package protos

type StickerLobbyRequest struct {
	RequestPacket
	AcquireStickerUniqueIds []int64 `json:",omitempty,omitzero"`
}
