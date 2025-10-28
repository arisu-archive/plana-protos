package protos

type EquipmentItemLockResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentDB EquipmentDB `json:",omitempty,omitzero"`
}
