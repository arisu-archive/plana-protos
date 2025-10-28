package protos

type ClanLoginResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountClanDB ClanDB `json:",omitempty,omitzero"`
	AccountClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	ClanAssistSlotDBs []ClanAssistSlotDB `json:",omitempty,omitzero"`
}
