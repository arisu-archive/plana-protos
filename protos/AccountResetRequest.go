package protos

type AccountResetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	DevId string `json:",omitempty,omitzero"`
}
