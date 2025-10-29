package protos

type EventContentBoxGachaShopListResponse struct {
	ResponsePacket
	BoxGachaDB EventContentBoxGachaDB `json:",omitempty,omitzero"`
	BoxGachaGroupIdByCount map[int64]int64 `json:",omitempty,omitzero"`
}
