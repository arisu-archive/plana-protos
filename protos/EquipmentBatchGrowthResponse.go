package protos

type EquipmentBatchGrowthResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentDBs []EquipmentDB `json:",omitempty,omitzero"`
	GearDB GearDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
