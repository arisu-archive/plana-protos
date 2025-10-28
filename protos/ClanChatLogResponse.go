package protos

type ClanChatLogResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanChatLog string `json:",omitempty,omitzero"`
}
