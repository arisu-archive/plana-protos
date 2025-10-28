package protos

type CafeUpdatePresetFurnitureRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	SlotId int32 `json:",omitempty,omitzero"`
}
