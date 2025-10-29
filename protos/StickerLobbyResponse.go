package protos

type StickerLobbyResponse struct {
	ResponsePacket
	ReceivedStickerDBs []StickerDB `json:",omitempty,omitzero"`
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
}
