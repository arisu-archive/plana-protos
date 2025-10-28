package protos

type StickerLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ReceivedStickerDBs []StickerDB `json:",omitempty,omitzero"`
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
}
