package protos

type CafeRemoveAllFurnitureRequest struct {
	RequestPacket
	CafeDBId int64 `json:",omitempty,omitzero"`
}
