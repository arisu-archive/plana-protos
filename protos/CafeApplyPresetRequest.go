package protos

type CafeApplyPresetRequest struct {
	RequestPacket
	SlotId int32 `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	UseOtherCafeFurniture bool `json:",omitempty,omitzero"`
}
