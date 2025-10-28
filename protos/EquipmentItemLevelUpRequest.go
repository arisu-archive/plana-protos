package protos

type EquipmentItemLevelUpRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetServerId int64 `json:",omitempty,omitzero"`
	ConsumeServerIds []int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
