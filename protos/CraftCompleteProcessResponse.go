package protos

type CraftCompleteProcessResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	CraftInfoDB       CraftInfoDB
	TicketItemDB      ItemDB
}
