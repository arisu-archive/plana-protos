package protos

type ItemSellRequest struct {
	RequestPacket
	TargetServerIds []int64
}
