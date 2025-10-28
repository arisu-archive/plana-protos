package protos

type EquipmentItemListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
