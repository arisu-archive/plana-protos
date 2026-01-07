package protos

type CafeUpdateCopyPresetFurnitureRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
	TargetCafeDBId  int64 `json:",omitempty,omitzero"`
	SlotId          int32 `json:",omitempty,omitzero"`
}
