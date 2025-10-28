package protos

type CraftCompleteProcessResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	CraftInfoDB CraftInfoDB `json:",omitempty,omitzero"`
	TicketItemDB ItemDB `json:",omitempty,omitzero"`
}
