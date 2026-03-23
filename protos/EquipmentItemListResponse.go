package protos

type EquipmentItemListResponse struct {
	ResponsePacket
	EquipmentDBs []*EquipmentDB
}
