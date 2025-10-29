package protos

type ItemSellResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
