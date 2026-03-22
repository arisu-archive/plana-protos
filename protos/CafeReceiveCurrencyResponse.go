package protos

type CafeReceiveCurrencyResponse struct {
	ResponsePacket
	CafeDB         CafeDB
	CafeDBs        []CafeDB
	ParcelResultDB ParcelResultDB
}
