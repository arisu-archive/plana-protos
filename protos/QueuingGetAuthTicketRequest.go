package protos

type QueuingGetAuthTicketRequest struct {
	RequestPacket
	ClientGeneratedKey string `json:",omitempty,omitzero"`
	ClientGeneratedIV  string `json:",omitempty,omitzero"`
	YostarUID          int64  `json:",omitempty,omitzero"`
	YostarToken        string `json:",omitempty,omitzero"`
	PassCheck          bool   `json:",omitempty,omitzero"`
	MakeStandby        bool   `json:",omitempty,omitzero"`
	PassCheckYostar    bool   `json:",omitempty,omitzero"`
	ClientVersion      string `json:",omitempty,omitzero"`
	OSType             string `json:",omitempty,omitzero"`
}
