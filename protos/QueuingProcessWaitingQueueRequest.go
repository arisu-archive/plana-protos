package protos

type QueuingProcessWaitingQueueRequest struct {
	RequestPacket
	WaitingTicket string
	ClientVersion string
	OSType        string
	AuthTicket    string
}
