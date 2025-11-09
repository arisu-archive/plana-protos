package protos

type EquipmentBatchGrowthRequest struct {
	RequestPacket
	EquipmentBatchGrowthRequestDBs []EquipmentBatchGrowthRequestDB `json:",omitempty,omitzero"`
	GearTierUpRequestDB            GearTierUpRequestDB             `json:",omitempty,omitzero"`
}
