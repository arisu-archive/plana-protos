package protos

type EventContentCardShopPurchaseResponse struct {
	ResponsePacket
	ParcelResultDB             ParcelResultDB              `json:",omitempty,omitzero"`
	CardShopElementDB          CardShopElementDB           `json:",omitempty,omitzero"`
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB `json:",omitempty,omitzero"`
}
