package protos

type EquipmentItemLockResponse struct {
	ResponsePacket
	EquipmentDB EquipmentDB `json:",omitempty,omitzero"`
}
