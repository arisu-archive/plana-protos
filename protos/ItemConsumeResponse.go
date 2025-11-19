package protos

type ItemConsumeResponse struct {
	ResponsePacket
	UsedItemDB        ItemDB         `json:",omitempty,omitzero"`
	NewParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
