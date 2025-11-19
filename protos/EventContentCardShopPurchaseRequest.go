package protos

type EventContentCardShopPurchaseRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	SlotNumber     int32 `json:",omitempty,omitzero"`
}
