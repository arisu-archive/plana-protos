package protos

type StickerUseStickerRequest struct {
	RequestPacket
	StickerUniqueId int64 `json:",omitempty,omitzero"`
}
