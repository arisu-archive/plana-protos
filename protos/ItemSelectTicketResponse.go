package protos

type ItemSelectTicketResponse struct {
	ResponsePacket
	UsedItemDB     ItemDB
	ParcelResultDB ParcelResultDB
}
