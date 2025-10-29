package protos

type ClanMemberResponse struct {
	ResponsePacket
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	DetailedAccountInfoDB DetailedAccountInfoDB `json:",omitempty,omitzero"`
}
