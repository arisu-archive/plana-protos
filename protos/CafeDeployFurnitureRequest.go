package protos

type CafeDeployFurnitureRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	FurnitureDB FurnitureDB `json:",omitempty,omitzero"`
}
