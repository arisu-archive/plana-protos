package protos

type ClanMyAssistListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanAssistSlotDBs []ClanAssistSlotDB `json:",omitempty,omitzero"`
}
