package protos

type EventContentCardShopListResponse struct {
	ResponsePacket
	CardShopElementDBs []CardShopElementDB `json:",omitempty,omitzero"`
	RewardHistory map[int64][]ParcelInfo `json:",omitempty,omitzero"`
}
