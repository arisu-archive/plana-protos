package protos

type ClanJoinResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IrcConfig IrcServerConfig `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
}
