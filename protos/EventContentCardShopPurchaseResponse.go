package protos

type EventContentCardShopPurchaseResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CardShopElementDB CardShopElementDB `json:",omitempty,omitzero"`
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB `json:",omitempty,omitzero"`
}
