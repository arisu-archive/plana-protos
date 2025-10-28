package protos

type ToastListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
