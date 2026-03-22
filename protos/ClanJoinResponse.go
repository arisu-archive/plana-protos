package protos

type ClanJoinResponse struct {
	ResponsePacket
	IrcConfig    IrcServerConfig
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
