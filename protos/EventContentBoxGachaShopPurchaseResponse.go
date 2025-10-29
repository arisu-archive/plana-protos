package protos

type EventContentBoxGachaShopPurchaseResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	BoxGachaDB EventContentBoxGachaDB `json:",omitempty,omitzero"`
	BoxGachaGroupIdByCount map[int64]int64 `json:",omitempty,omitzero"`
	BoxGachaElements []EventContentBoxGachaElement `json:",omitempty,omitzero"`
}
