package protos

type EquipmentItemTierUpRequest struct {
	RequestPacket
	TargetEquipmentServerId int64                     `json:",omitempty,omitzero"`
	ReplaceInfos            []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
