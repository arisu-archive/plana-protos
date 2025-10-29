package protos

type ItemSelectTicketResponse struct {
	ResponsePacket
	UsedItemDB ItemDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
