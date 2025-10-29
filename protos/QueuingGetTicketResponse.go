package protos

type QueuingGetTicketResponse struct {
	ResponsePacket
	WaitingTicket string `json:",omitempty,omitzero"`
	EnterTicket string `json:",omitempty,omitzero"`
	TicketSequence int64 `json:",omitempty,omitzero"`
	AllowedSequence int64 `json:",omitempty,omitzero"`
	RequiredSecondsPerUser float64 `json:",omitempty,omitzero"`
	Birth string `json:",omitempty,omitzero"`
	ServerSeed string `json:",omitempty,omitzero"`
}
