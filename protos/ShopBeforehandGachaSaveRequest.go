package protos

type ShopBeforehandGachaSaveRequest struct {
	RequestPacket
	TargetIndex int64  `json:",omitempty,omitzero"`
	SlotNumber  *int64 `json:",omitempty,omitzero"`
}
