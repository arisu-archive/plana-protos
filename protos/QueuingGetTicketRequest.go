package protos

type QueuingGetTicketRequest struct {
	RequestPacket
	YostarUID int64 `json:",omitempty,omitzero"`
	YostarToken string `json:",omitempty,omitzero"`
	MakeStandby bool `json:",omitempty,omitzero"`
	PassCheck bool `json:",omitempty,omitzero"`
	PassCheckYostar bool `json:",omitempty,omitzero"`
	WaitingTicket string `json:",omitempty,omitzero"`
	ClientVersion string `json:",omitempty,omitzero"`
	OSType string `json:",omitempty,omitzero"`
}
