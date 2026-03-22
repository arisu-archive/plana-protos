package protos

type CraftInfoListResponse struct {
	ResponsePacket
	CraftInfos         []CraftInfoDB
	ShiftingCraftInfos []ShiftingCraftInfoDB
	PresetSlotDBs      []CraftPresetSlotDB
}
