package protos

type ItemConsumeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UsedItemDB ItemDB `json:",omitempty,omitzero"`
	NewParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
