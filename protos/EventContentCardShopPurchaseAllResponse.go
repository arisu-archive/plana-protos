package protos

type EventContentCardShopPurchaseAllResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CardShopElementDBs []CardShopElementDB `json:",omitempty,omitzero"`
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB `json:",omitempty,omitzero"`
	RewardHistory map[int64][]ParcelInfo `json:",omitempty,omitzero"`
}
