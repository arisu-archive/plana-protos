package protos

type EquipmentBatchGrowthResponse struct {
	ResponsePacket
	EquipmentDBs    []*EquipmentDB
	GearDB          *GearDB          `json:",omitempty,omitzero"`
	ParcelResultDB  *ParcelResultDB  `json:",omitempty,omitzero"`
	ConsumeResultDB *ConsumeResultDB `json:",omitempty,omitzero"`
}
