package protos

type AccountDismissRepurchasablePopupRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ProductIds []int64 `json:",omitempty,omitzero"`
}
