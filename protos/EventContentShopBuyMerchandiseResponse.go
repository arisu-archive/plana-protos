package protos

type EventContentShopBuyMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB         AccountCurrencyDB
	ConsumeResultDB           ConsumeResultDB
	ParcelResultDB            ParcelResultDB
	MailDB                    MailDB
	ShopProductDB             ShopProductDB
	EventContentCollectionDBs []EventContentCollectionDB
}
