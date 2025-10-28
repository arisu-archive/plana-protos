package protos

type ClanMemberResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	DetailedAccountInfoDB DetailedAccountInfoDB `json:",omitempty,omitzero"`
}
