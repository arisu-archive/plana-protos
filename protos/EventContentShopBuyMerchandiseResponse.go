package protos

type EventContentShopBuyMerchandiseResponse struct {
	ResponsePacket
	AccountCurrencyDB         AccountCurrencyDB          `json:",omitempty,omitzero"`
	ConsumeResultDB           ConsumeResultDB            `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB             `json:",omitempty,omitzero"`
	MailDB                    MailDB                     `json:",omitempty,omitzero"`
	ShopProductDB             ShopProductDB              `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
