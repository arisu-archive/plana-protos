package protos

type CafeRemoveFurnitureRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
	FurnitureServerIds []int64 `json:",omitempty,omitzero"`
}
