package protos

type ShopBuyAPResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	MailDB            MailDB
	ShopProductDB     ShopProductDB
}
