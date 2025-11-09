package protos

type CafeRemoveAllFurnitureResponse struct {
	ResponsePacket
	CafeDB       CafeDB        `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
