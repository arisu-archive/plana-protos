package protos

type CafeSummonCharacterTicketUseResponse struct {
	ResponsePacket
	CafeDB         CafeDB
	CafeDBs        []CafeDB
	ParcelResultDB ParcelResultDB
}
