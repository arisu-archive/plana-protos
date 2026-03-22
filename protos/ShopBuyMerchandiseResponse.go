package protos

type ShopBuyMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	MailDB            MailDB
	ShopProductDB     ShopProductDB
}
