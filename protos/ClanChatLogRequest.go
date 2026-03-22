package protos

type ClanChatLogRequest struct {
	RequestPacket
	Channel  string
	FromDate MxTime `json:",omitempty,omitzero"`
}
