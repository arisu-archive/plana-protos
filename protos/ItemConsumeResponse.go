package protos

type ItemConsumeResponse struct {
	ResponsePacket
	UsedItemDB        ItemDB
	NewParcelResultDB ParcelResultDB
}
