package protos

type ItemBulkConsumeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetItemServerId int64 `json:",omitempty,omitzero"`
	ConsumeCount int32 `json:",omitempty,omitzero"`
}
