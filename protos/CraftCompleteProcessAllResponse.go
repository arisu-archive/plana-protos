package protos

type CraftCompleteProcessAllResponse struct {
	ResponsePacket
	CraftInfoDBs []CraftInfoDB `json:",omitempty,omitzero"`
	TicketItemDB ItemDB `json:",omitempty,omitzero"`
}
