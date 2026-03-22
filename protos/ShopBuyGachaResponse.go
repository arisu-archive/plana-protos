package protos

type ShopBuyGachaResponse struct {
	ResponsePacket
	ConsumeResultDB ConsumeResultDB
	ParcelResultDB  ParcelResultDB
}
