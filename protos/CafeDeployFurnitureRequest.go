package protos

type CafeDeployFurnitureRequest struct {
	RequestPacket
	CafeDBId    int64       `json:",omitempty,omitzero"`
	FurnitureDB FurnitureDB `json:",omitempty,omitzero"`
}
