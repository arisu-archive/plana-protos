package protos

type EventContentCardShopPurchaseResponse struct {
	ResponsePacket
	ParcelResultDB             ParcelResultDB
	CardShopElementDB          CardShopElementDB
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB
}
