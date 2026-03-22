package protos

type ShopBuyEligmaResponse struct {
	ResponsePacket
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
	ShopProductDB   ShopProductDB
}
