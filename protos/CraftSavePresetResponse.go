package protos

type CraftSavePresetResponse struct {
	ResponsePacket
	PresetSlotDB CraftPresetSlotDB `json:",omitempty,omitzero"`
}
