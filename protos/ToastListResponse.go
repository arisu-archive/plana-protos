package protos

type ToastListResponse struct {
	ResponsePacket
	ToastDBs []ToastDB `json:",omitempty,omitzero"`
}
