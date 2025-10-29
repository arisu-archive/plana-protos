package protos

type AccountDismissRepurchasablePopupRequest struct {
	RequestPacket
	ProductIds []int64 `json:",omitempty,omitzero"`
}
