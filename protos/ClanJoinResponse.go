package protos

type ClanJoinResponse struct {
	ResponsePacket
	IrcConfig    IrcServerConfig `json:",omitempty,omitzero"`
	ClanDB       ClanDB          `json:",omitempty,omitzero"`
	ClanMemberDB ClanMemberDB    `json:",omitempty,omitzero"`
}
