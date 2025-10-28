package protos

type StickerUseStickerRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StickerUniqueId int64 `json:",omitempty,omitzero"`
}
