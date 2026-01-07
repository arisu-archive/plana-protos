package protos

type AccountCurrencySyncResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ExpiredCurrency   map[string]int64  `json:",omitempty,omitzero"`
}
