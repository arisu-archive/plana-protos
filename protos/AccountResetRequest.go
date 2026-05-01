package protos

type AccountResetRequest struct {
	RequestPacket
	DevId string `json:",omitempty,omitzero"`
}
