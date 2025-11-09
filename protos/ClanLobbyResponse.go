package protos

type ClanLobbyResponse struct {
	ResponsePacket
	IrcConfig             IrcServerConfig `json:",omitempty,omitzero"`
	AccountClanDB         ClanDB          `json:",omitempty,omitzero"`
	DefaultExposedClanDBs []ClanDB        `json:",omitempty,omitzero"`
	AccountClanMemberDB   ClanMemberDB    `json:",omitempty,omitzero"`
	ClanMemberDBs         []ClanMemberDB  `json:",omitempty,omitzero"`
}
