package protos

type CafeRemoveFurnitureResponse struct {
	ResponsePacket
	CafeDB       CafeDB
	FurnitureDBs []FurnitureDB
}
