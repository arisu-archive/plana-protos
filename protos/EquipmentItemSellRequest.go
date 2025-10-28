package protos

type EquipmentItemSellRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetServerIds []int64 `json:",omitempty,omitzero"`
}
