package protos

type CafeRemoveFurnitureRequest struct {
	RequestPacket
	CafeDBId           int64   `json:",omitempty,omitzero"`
	FurnitureServerIds []int64 `json:",omitempty,omitzero"`
}
