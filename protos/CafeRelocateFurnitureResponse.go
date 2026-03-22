package protos

type CafeRelocateFurnitureResponse struct {
	ResponsePacket
	CafeDB               CafeDB
	RelocatedFurnitureDB FurnitureDB
}
