package protos

type CafeRemoveAllFurnitureResponse struct {
	ResponsePacket
	CafeDB       CafeDB
	FurnitureDBs []FurnitureDB
}
