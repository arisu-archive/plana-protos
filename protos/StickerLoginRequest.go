package protos

type StickerLoginRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
