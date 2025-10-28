package protos

type CafeOpenRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeId int64 `json:",omitempty,omitzero"`
}
