package protos

type CraftAutoBeginProcessRequest struct {
	RequestPacket
	PresetSlotDB CraftPresetSlotDB `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
