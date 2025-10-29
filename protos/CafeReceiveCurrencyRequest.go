package protos

type CafeReceiveCurrencyRequest struct {
	RequestPacket
	AccountServerId int64 `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
}
