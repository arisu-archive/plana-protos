package protos

type WorldRaidLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClearHistoryDBs []WorldRaidClearHistoryDB `json:",omitempty,omitzero"`
	LocalBossDBs []WorldRaidLocalBossDB `json:",omitempty,omitzero"`
	BossGroups []WorldRaidBossGroup `json:",omitempty,omitzero"`
}
