package protos

type ClanAutoJoinResponse struct {
	ResponsePacket
	IrcConfig    IrcServerConfig
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
