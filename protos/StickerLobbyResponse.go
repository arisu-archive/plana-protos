package protos

type StickerLobbyResponse struct {
	ResponsePacket
	ReceivedStickerDBs []*StickerDB
	StickerBookDB      *StickerBookDB `json:",omitempty,omitzero"`
}
