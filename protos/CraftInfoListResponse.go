package protos

type CraftInfoListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfos []CraftInfoDB `json:",omitempty,omitzero"`
	ShiftingCraftInfos []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
}
