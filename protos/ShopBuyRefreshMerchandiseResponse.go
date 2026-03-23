package protos

type ShopBuyRefreshMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB *AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB   *ConsumeResultDB   `json:",omitempty,omitzero"`
	ParcelResultDB    *ParcelResultDB    `json:",omitempty,omitzero"`
	ShopProductDB     []*ShopProductDB
	MailDB            *MailDB `json:",omitempty,omitzero"`
}
