package protos

type ClanLobbyResponse struct {
	ResponsePacket
	IrcConfig             *IrcServerConfig `json:",omitempty,omitzero"`
	AccountClanDB         *ClanDB          `json:",omitempty,omitzero"`
	DefaultExposedClanDBs []ClanDB
	AccountClanMemberDB   *ClanMemberDB `json:",omitempty,omitzero"`
	ClanMemberDBs         []ClanMemberDB
}
