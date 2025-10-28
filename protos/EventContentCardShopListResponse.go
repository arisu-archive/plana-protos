package protos

type EventContentCardShopListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CardShopElementDBs []CardShopElementDB `json:",omitempty,omitzero"`
	RewardHistory map[int64][]ParcelInfo `json:",omitempty,omitzero"`
}
