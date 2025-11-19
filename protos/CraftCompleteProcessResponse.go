package protos

type CraftCompleteProcessResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	CraftInfoDB       CraftInfoDB       `json:",omitempty,omitzero"`
	TicketItemDB      ItemDB            `json:",omitempty,omitzero"`
}
