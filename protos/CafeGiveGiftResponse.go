package protos

type CafeGiveGiftResponse struct {
	ResponsePacket
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
