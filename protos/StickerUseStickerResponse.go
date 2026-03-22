package protos

type StickerUseStickerResponse struct {
	ResponsePacket
	StickerBookDB  StickerBookDB
	ParcelResultDB ParcelResultDB
}
