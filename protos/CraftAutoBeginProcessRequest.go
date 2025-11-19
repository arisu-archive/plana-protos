package protos

type CraftAutoBeginProcessRequest struct {
	RequestPacket
	PresetIndex int32 `json:",omitempty,omitzero"`
	Count       int64 `json:",omitempty,omitzero"`
}
