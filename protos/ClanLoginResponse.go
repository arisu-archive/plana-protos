package protos

type ClanLoginResponse struct {
	ResponsePacket
	AccountClanDB ClanDB `json:",omitempty,omitzero"`
	AccountClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	ClanAssistSlotDBs []ClanAssistSlotDB `json:",omitempty,omitzero"`
}
