package protos

type CafeGetInfoResponse struct {
	ResponsePacket
	CafeDB       CafeDB
	CafeDBs      []CafeDB
	FurnitureDBs []FurnitureDB
}
