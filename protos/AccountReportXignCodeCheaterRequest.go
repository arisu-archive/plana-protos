package protos

type AccountReportXignCodeCheaterRequest struct {
	RequestPacket
	ErrorCode string `json:",omitempty,omitzero"`
}
