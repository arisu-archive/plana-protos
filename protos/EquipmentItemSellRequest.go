package protos

type EquipmentItemSellRequest struct {
	RequestPacket
	TargetServerIds []int64
}
