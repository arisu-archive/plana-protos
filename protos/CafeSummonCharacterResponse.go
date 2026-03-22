package protos

type CafeSummonCharacterResponse struct {
	ResponsePacket
	CafeDB  CafeDB
	CafeDBs []CafeDB
}
