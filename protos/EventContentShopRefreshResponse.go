package protos

type EventContentShopRefreshResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	ShopInfoDB     ShopInfoDB
}
