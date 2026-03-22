package protos

type CraftCompleteProcessAllResponse struct {
	ResponsePacket
	CraftInfoDBs []CraftInfoDB
	TicketItemDB ItemDB
}
