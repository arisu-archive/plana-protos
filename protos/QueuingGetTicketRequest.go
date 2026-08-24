package protos

type QueuingGetTicketRequest struct {
	RequestPacket
	YostarUID       int64 `json:",omitempty,omitzero"`
	YostarToken     string
	MakeStandby     bool `json:",omitempty,omitzero"`
	PassCheck       bool `json:",omitempty,omitzero"`
	PassCheckYostar bool `json:",omitempty,omitzero"`
	WaitingTicket   string
	ClientVersion   string
	OSType          string
}
