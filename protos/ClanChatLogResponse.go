package protos

type ClanChatLogResponse struct {
	ResponsePacket
	ClanChatLog string `json:",omitempty,omitzero"`
}
