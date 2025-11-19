package protos

type CraftInfoListResponse struct {
	ResponsePacket
	CraftInfos         []CraftInfoDB         `json:",omitempty,omitzero"`
	ShiftingCraftInfos []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	PresetSlotDBs      []CraftPresetSlotDB   `json:",omitempty,omitzero"`
}
