package protos

type EventContentShopBuyRefreshMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB         AccountCurrencyDB
	ConsumeResultDB           ConsumeResultDB
	ParcelResultDB            ParcelResultDB
	MailDB                    MailDB
	ShopProductDB             []ShopProductDB
	EventContentCollectionDBs []EventContentCollectionDB
}
