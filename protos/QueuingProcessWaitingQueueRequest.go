package protos

type QueuingProcessWaitingQueueRequest struct {
	RequestPacket
	WaitingTicket string `json:",omitempty,omitzero"`
	ClientVersion string `json:",omitempty,omitzero"`
	OSType        string `json:",omitempty,omitzero"`
	AuthTicket    string `json:",omitempty,omitzero"`
}
