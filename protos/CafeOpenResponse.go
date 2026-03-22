package protos

type CafeOpenResponse struct {
	ResponsePacket
	OpenedCafeDB CafeDB
	FurnitureDBs []FurnitureDB
}
