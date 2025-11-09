package protos

type CafeUpdatePresetFurnitureRequest struct {
	RequestPacket
	CafeDBId int64 `json:",omitempty,omitzero"`
	SlotId   int32 `json:",omitempty,omitzero"`
}
