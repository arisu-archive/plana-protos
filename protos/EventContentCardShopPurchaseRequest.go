package protos

type EventContentCardShopPurchaseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	SlotNumber int32 `json:",omitempty,omitzero"`
}
