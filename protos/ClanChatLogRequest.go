package protos

type ClanChatLogRequest struct {
	RequestPacket
	Channel  string `json:",omitempty,omitzero"`
	FromDate MxTime `json:",omitempty,omitzero"`
}
