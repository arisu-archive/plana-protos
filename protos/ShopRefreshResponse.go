package protos

type ShopRefreshResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	ShopInfoDB     ShopInfoDB
}
