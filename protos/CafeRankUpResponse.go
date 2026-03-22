package protos

type CafeRankUpResponse struct {
	ResponsePacket
	CafeDB          CafeDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
