package protos

type ClanLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IrcConfig IrcServerConfig `json:",omitempty,omitzero"`
	AccountClanDB ClanDB `json:",omitempty,omitzero"`
	DefaultExposedClanDBs []ClanDB `json:",omitempty,omitzero"`
	AccountClanMemberDB ClanMemberDB `json:",omitempty,omitzero"`
	ClanMemberDBs []ClanMemberDB `json:",omitempty,omitzero"`
}
