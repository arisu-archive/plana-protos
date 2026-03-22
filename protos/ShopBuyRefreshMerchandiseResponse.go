package protos

type ShopBuyRefreshMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	ShopProductDB     []ShopProductDB
	MailDB            MailDB
}
