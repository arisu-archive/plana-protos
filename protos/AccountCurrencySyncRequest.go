package protos

type AccountCurrencySyncRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
