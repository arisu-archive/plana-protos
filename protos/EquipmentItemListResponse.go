package protos

type EquipmentItemListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EquipmentDBs []EquipmentDB `json:",omitempty,omitzero"`
}
