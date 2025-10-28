package protos

type CafeRelocateFurnitureResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDB CafeDB `json:",omitempty,omitzero"`
	RelocatedFurnitureDB FurnitureDB `json:",omitempty,omitzero"`
}
