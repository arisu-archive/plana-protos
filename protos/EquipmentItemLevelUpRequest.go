package protos

type EquipmentItemLevelUpRequest struct {
	RequestPacket
	TargetServerId int64 `json:",omitempty,omitzero"`
	ConsumeServerIds []int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
