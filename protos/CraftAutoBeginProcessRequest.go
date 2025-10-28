package protos

type CraftAutoBeginProcessRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetSlotDB CraftPresetSlotDB `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
