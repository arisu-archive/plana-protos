package protos

type QueuingGetTicketResponse struct {
	ResponsePacket
	WaitingTicket          string
	EnterTicket            string
	TicketSequence         int64   `json:",omitempty,omitzero"`
	AllowedSequence        int64   `json:",omitempty,omitzero"`
	RequiredSecondsPerUser float64 `json:",omitempty,omitzero"`
	Birth                  string
	ServerSeed             string
}
