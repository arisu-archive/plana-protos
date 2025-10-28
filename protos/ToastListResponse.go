package protos

type ToastListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ToastDBs []ToastDB `json:",omitempty,omitzero"`
}
