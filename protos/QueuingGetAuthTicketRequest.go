package protos

type QueuingGetAuthTicketRequest struct {
	RequestPacket
	ClientGeneratedKey string
	ClientGeneratedIV  string
	YostarUID          int64 `json:",omitempty,omitzero"`
	YostarToken        string
	PassCheck          bool `json:",omitempty,omitzero"`
	MakeStandby        bool `json:",omitempty,omitzero"`
	PassCheckYostar    bool `json:",omitempty,omitzero"`
	ClientVersion      string
	OSType             string
}
