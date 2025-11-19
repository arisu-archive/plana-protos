package protos

type CafeRemoveFurnitureResponse struct {
	ResponsePacket
	CafeDB       CafeDB        `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
}
