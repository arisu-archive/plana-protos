package protos

type StickerUseStickerResponse struct {
	ResponsePacket
	StickerBookDB StickerBookDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
