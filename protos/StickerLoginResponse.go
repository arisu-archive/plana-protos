package protos

type StickerLoginResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
}
