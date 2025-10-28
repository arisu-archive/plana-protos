package protos

type EquipmentItemTierUpRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetEquipmentServerId int64 `json:",omitempty,omitzero"`
	ReplaceInfos []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
