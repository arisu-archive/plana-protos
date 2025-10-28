package protos

type ItemSelectTicketResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UsedItemDB ItemDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
