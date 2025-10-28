package protos

type ClanMyAssistListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
