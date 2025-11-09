package protos

type ClanConferResponse struct {
	ResponsePacket
	ClanMemberDB            ClanMemberDB            `json:",omitempty,omitzero"`
	AccountClanMemberDB     ClanMemberDB            `json:",omitempty,omitzero"`
	ClanDB                  ClanDB                  `json:",omitempty,omitzero"`
	ClanMemberDescriptionDB ClanMemberDescriptionDB `json:",omitempty,omitzero"`
}
