package protos

type EquipmentBatchGrowthRequest struct {
	RequestPacket
	EquipmentBatchGrowthRequestDBs []EquipmentBatchGrowthRequestDB
	GearTierUpRequestDB            *GearTierUpRequestDB `json:",omitempty,omitzero"`
}
