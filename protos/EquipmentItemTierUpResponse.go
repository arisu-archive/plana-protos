package protos

type EquipmentItemTierUpResponse struct {
	ResponsePacket
	EquipmentDB     EquipmentDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
