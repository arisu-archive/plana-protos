package protos

type StickerLoginResponse struct {
	ResponsePacket
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
}
