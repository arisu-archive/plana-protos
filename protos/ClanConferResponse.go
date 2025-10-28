package protos

type ClanConferResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	AccountClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDescriptionDB ClanMemberDescriptionDB `json:",omitempty,omitzero"`
}
