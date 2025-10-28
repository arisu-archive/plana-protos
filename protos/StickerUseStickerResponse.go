package protos

type StickerUseStickerResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
