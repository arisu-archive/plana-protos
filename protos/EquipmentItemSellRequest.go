package protos

type EquipmentItemSellRequest struct {
	RequestPacket
	TargetServerIds []int64 `json:",omitempty,omitzero"`
}
