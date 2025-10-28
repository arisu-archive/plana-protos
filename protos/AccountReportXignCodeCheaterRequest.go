package protos

type AccountReportXignCodeCheaterRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ErrorCode string `json:",omitempty,omitzero"`
}
