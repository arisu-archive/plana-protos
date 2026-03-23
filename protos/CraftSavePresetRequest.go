package protos

type CraftSavePresetRequest struct {
	RequestPacket
	PresetSlotDB *CraftPresetSlotDB `json:",omitempty,omitzero"`
}
