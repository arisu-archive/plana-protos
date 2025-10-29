package protos

type ItemBulkConsumeRequest struct {
	RequestPacket
	TargetItemServerId int64 `json:",omitempty,omitzero"`
	ConsumeCount int32 `json:",omitempty,omitzero"`
}
