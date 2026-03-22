package protos

type ClanLobbyResponse struct {
	ResponsePacket
	IrcConfig             IrcServerConfig
	AccountClanDB         ClanDB
	DefaultExposedClanDBs []ClanDB
	AccountClanMemberDB   ClanMemberDB
	ClanMemberDBs         []ClanMemberDB
}
