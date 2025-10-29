package protos

type CafeOpenRequest struct {
	RequestPacket
	CafeId int64 `json:",omitempty,omitzero"`
}
