package protos

type ShopBeforehandGachaGetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
