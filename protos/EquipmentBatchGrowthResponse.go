package protos

type EquipmentBatchGrowthResponse struct {
	ResponsePacket
	EquipmentDBs    []EquipmentDB
	GearDB          GearDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
