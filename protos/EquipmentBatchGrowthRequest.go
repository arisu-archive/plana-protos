package protos

type EquipmentBatchGrowthRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentBatchGrowthRequestDBs []EquipmentBatchGrowthRequestDB `json:",omitempty,omitzero"`
	GearTierUpRequestDB GearTierUpRequestDB `json:",omitempty,omitzero"`
}
