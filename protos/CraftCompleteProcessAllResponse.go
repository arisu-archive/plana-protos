package protos

type CraftCompleteProcessAllResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfoDBs []CraftInfoDB `json:",omitempty,omitzero"`
	TicketItemDB ItemDB `json:",omitempty,omitzero"`
}
