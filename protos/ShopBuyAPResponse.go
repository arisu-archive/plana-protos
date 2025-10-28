package protos

type ShopBuyAPResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	MailDB MailDB `json:",omitempty,omitzero"`
	ShopProductDB ShopProductDB `json:",omitempty,omitzero"`
}
