package protos

type ShopBeforehandGachaSaveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetIndex int64 `json:",omitempty,omitzero"`
}
