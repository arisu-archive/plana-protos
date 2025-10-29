package protos

type ClanMyAssistListResponse struct {
	ResponsePacket
	ClanAssistSlotDBs []ClanAssistSlotDB `json:",omitempty,omitzero"`
}
