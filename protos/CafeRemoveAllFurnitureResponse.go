package protos

type CafeRemoveAllFurnitureResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDB CafeDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
