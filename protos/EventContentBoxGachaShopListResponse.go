package protos

type EventContentBoxGachaShopListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BoxGachaDB EventContentBoxGachaDB `json:",omitempty,omitzero"`
	BoxGachaGroupIdByCount map[int64]int64 `json:",omitempty,omitzero"`
}
