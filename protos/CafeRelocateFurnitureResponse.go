package protos

type CafeRelocateFurnitureResponse struct {
	ResponsePacket
	CafeDB CafeDB `json:",omitempty,omitzero"`
	RelocatedFurnitureDB FurnitureDB `json:",omitempty,omitzero"`
}
